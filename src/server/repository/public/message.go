package repository

import (
	"context"
	localization "earnforglance/server/domain/localization"
	message "earnforglance/server/domain/messages"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	seo "earnforglance/server/domain/seo"
	stores "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type newsLetterRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsLetterRepository(db mongo.Database, collection string) domain.NewsLetterRepository {
	return &newsLetterRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *newsLetterRepository) ContactUs(c context.Context, contact domain.ContactUsRequest) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	result, err = PrepareContactUs(cr, c, contact)
	if err != nil || !result.Result {
		return result, err
	}

	if contact.News {
		request := domain.NewsLetterRequest{Email: contact.Email, StoreID: contact.StoreID, Lang: contact.Lang}
		result, err = PrepareNewsLetterSubscription(cr, c, request, contact.IpAddress)
		if err != nil || !result.Result {
			return result, err
		}
	}

	return result, err
}

func PrepareContactUs(nr *newsLetterRepository, c context.Context, contact domain.ContactUsRequest) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	langID, err := GetLangugaByCode(c, contact.Lang, nr.database.Collection(localization.CollectionLanguage))
	if err != nil {
		result.Result = false
		result.Message = "Invalid Language Code"
		return result, err
	}

	if contact.Email == "" {
		result.Result = false
		result.Message = "Invalid email"
		return result, err
	}

	emailAccount, err := GetEmailAcount(c, nr.database.Collection(message.CollectionEmailAccount))
	if err != nil || emailAccount == nil {
		result.Result = false
		result.Message = "Failed to get email account"
		return result, err
	}

	// Send email to the user for confirmation to deactivate the subscription
	template, err := GetMessageTemplateEmail(nr, c, "CONTACT_US_MESSAGE", langID.ID, nr.database.Collection(message.CollectionMessageTemplate))
	if err != nil {
		return result, err
	}

	for _, store := range contact.StoreID {

		storeID, err := bson.ObjectIDFromHex(store)
		if err != nil {
			result.Result = false
			result.Message = "Invalid Store ID"
			return result, err
		}

		fields, err := GetFieldsByID(c, storeID, nr.database.Collection(stores.CollectionStore))
		if err != nil {
			result.Result = false
			result.Message = "Failed to get store fields"
			return result, err
		}

		tokensSubject := GetTokens(c, template.Subject)
		template.Subject = ReplaceTokens(c, template.Subject, tokensSubject, fields, stores.CollectionStore)

		tokensBody := GetTokens(c, template.Body)
		template.Body = ReplaceTokens(c, template.Body, tokensBody, fields, stores.CollectionStore)

		// Convert suscriptionNews (struct) to bson.M
		doc, err := bson.Marshal(contact)
		if err != nil {
			result.Result = false
			result.Message = "Failed to convert suscriptionNews (struct) to bson.M"
			return result, err
		}
		// Convert doc to bson.M
		var bsonMap bson.M
		err = bson.Unmarshal(doc, &bsonMap)
		if err != nil {
			result.Result = false
			result.Message = "Failed to convert suscriptionNews doc to bson.M"
			return result, err
		}

		template.Body = ReplaceTokens(c, template.Body, tokensBody, &bsonMap, "ContactUs")

		ok, err := AddQueuedEmail(c, *emailAccount, emailAccount.Email, "", "", int(message.Low), template.Subject, template.Body, bson.ObjectID{}, nr.database.Collection(message.CollectionQueuedEmail))
		if err != nil || !ok {
			result.Result = false
			result.Message = "Failed to add queued email"
			return result, err
		}
	}

	locale, err := GetLocalebyName(c, "Newsletter.UnsubscribeEmailSent", langID.ID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
	result.Result = true
	result.Message = locale.ResourceValue
	return result, err
}

func (cr *newsLetterRepository) NewsLetterUnSubscribe(c context.Context, filter domain.NewsLetterRequest) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	for _, store := range filter.StoreID {

		storeID, err := bson.ObjectIDFromHex(store)
		if err != nil {
			result.Result = false
			result.Message = "Invalid Store ID"
			return result, err
		}

		result, err = PrepareNewsLetterUnSubscribe(cr, c, filter.Email, storeID, filter.Lang)

		if err != nil || !result.Result {
			return result, err
		}
	}

	return result, err
}

func PrepareNewsLetterUnSubscribe(nr *newsLetterRepository, c context.Context, email string, storeID bson.ObjectID, lang string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	if email == "" {
		result.Result = false
		result.Message = "Invalid email"
		return result, err
	}

	query := bson.M{
		"email":    email,
		"store_id": storeID,
	}

	suscriptionNews := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	collection.FindOne(c, query).Decode(&suscriptionNews)

	langID, err := GetLangugaByCode(c, lang, nr.database.Collection(localization.CollectionLanguage))
	if err != nil {
		result.Result = false
		result.Message = "Invalid Language Code"
		return result, err
	}

	emailAccount, err := GetEmailAcount(c, nr.database.Collection(message.CollectionEmailAccount))
	if err != nil || emailAccount == nil {
		result.Result = false
		result.Message = "Failed to get email account"
		return result, err
	}

	// Send email to the user for confirmation to deactivate the subscription
	template, err := GetMessageTemplateEmail(nr, c, "NEWSLETTER_SUBSCRIPTION_DEACTIVATION_MESSAGE", langID.ID, nr.database.Collection(message.CollectionMessageTemplate))
	if err != nil {
		return result, err
	}

	fields, err := GetFieldsByID(c, storeID, nr.database.Collection(stores.CollectionStore))
	if err != nil {
		result.Result = false
		result.Message = "Failed to get store fields"
		return result, err
	}

	tokensSubject := GetTokens(c, template.Subject)
	template.Subject = ReplaceTokens(c, template.Subject, tokensSubject, fields, stores.CollectionStore)

	slugs, err := GetSlugsNewsLetterbyLang(nr, c, message.CollectionNewsLetterSubscription, langID.ID)
	if err != nil {
		return result, err
	}

	tokensBody := GetTokens(c, template.Body)
	template.Body = ReplaceTokens(c, template.Body, tokensBody, fields, stores.CollectionStore)
	template.Body = ReplaceTokens(c, template.Body, tokensBody, slugs, seo.CollectionUrlRecord)

	// Convert suscriptionNews (struct) to bson.M
	doc, err := bson.Marshal(suscriptionNews)
	if err != nil {
		result.Result = false
		result.Message = "Failed to convert suscriptionNews (struct) to bson.M"
		return result, err
	}
	// Convert doc to bson.M
	var bsonMap bson.M
	err = bson.Unmarshal(doc, &bsonMap)
	if err != nil {
		result.Result = false
		result.Message = "Failed to convert suscriptionNews doc to bson.M"
		return result, err
	}

	template.Body = ReplaceTokens(c, template.Body, tokensBody, &bsonMap, message.CollectionNewsLetterSubscription)

	ok, err := AddQueuedEmail(c, *emailAccount, email, "", "", message.High, template.Subject, template.Body, bson.ObjectID{}, nr.database.Collection(message.CollectionQueuedEmail))
	if err != nil || !ok {
		result.Result = false
		result.Message = "Failed to add queued email"
		return result, err
	}

	locale, err := GetLocalebyName(c, "Newsletter.UnsubscribeEmailSent", langID.ID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
	result.Result = true
	result.Message = locale.ResourceValue
	return result, err
}

func (cr *newsLetterRepository) NewsLetterInactivate(c context.Context, Guid string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	result, err = PrepareNewsLetterInactivate(cr, c, Guid)

	if err != nil || !result.Result {
		return result, err
	}

	return result, err
}

func PrepareNewsLetterInactivate(nr *newsLetterRepository, c context.Context, guid string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	if guid == "" {
		result.Result = false
		result.Message = "Invalid Guid"
		return result, err
	}

	query := bson.M{
		"newsletter_subscription_guid": guid,
	}

	news := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	err = collection.FindOne(c, query).Decode(&news)
	if err != nil {
		result.Result = false
		result.Message = err.Error()
		return result, err
	}

	if news.Guid != guid {
		locale, err := GetLocalebyName(c, "Newsletter.ResultActivated.InvalidGuid", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		result.Result = false
		result.Message = locale.ResourceValue
		return result, err
	}

	if news.Active {
		// Deactivate the subscription
		update, err := collection.UpdateOne(c, query, bson.M{"$set": bson.M{"active": false, "newsletter_subscription_guid": uuid.New().String()}})
		if err != nil {
			result.Result = false
			result.Message = "Failed to update subscription"
			return result, err
		}
		if update.MatchedCount == 0 {
			result.Result = false
			result.Message = "No matching subscription found"
			return result, err
		}
	}

	locale, err := GetLocalebyName(c, "Newsletter.ResultDeactivated", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
	result.Result = true
	result.Message = locale.ResourceValue

	return result, err
}

func (cr *newsLetterRepository) NewsLetterActivation(c context.Context, Guid string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	result, err = PrepareNewsLetterActivation(cr, c, Guid)

	if err != nil || !result.Result {
		return result, err
	}

	return result, err
}

func PrepareNewsLetterActivation(nr *newsLetterRepository, c context.Context, guid string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	if guid == "" {
		result.Result = false
		result.Message = "Invalid Guid"
		return result, err
	}

	query := bson.M{
		"newsletter_subscription_guid": guid,
	}

	news := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	err = collection.FindOne(c, query).Decode(&news)
	if err != nil {
		result.Result = false
		result.Message = err.Error()
		return result, err
	}

	if news.Guid != guid {
		locale, err := GetLocalebyName(c, "Newsletter.ResultActivated.InvalidGuid", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		result.Result = false
		result.Message = locale.ResourceValue
		return result, err
	}

	if !news.Active {
		// Activate the subscription
		update, err := collection.UpdateOne(c, query, bson.M{"$set": bson.M{"active": true, "newsletter_subscription_guid": uuid.New().String()}})
		if err != nil {
			result.Result = false
			result.Message = "Failed to update subscription"
			return result, err
		}
		if update.MatchedCount == 0 {
			result.Result = false
			result.Message = "No matching subscription found"
			return result, err
		}
	}

	locale, err := GetLocalebyName(c, "Newsletter.ResultActivated", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
	result.Result = true
	result.Message = locale.ResourceValue

	return result, err
}

func (cr *newsLetterRepository) NewsLetterSubscription(c context.Context, filter domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)
	//mail.SendEmail()
	result, err = PrepareNewsLetterSubscription(cr, c, filter, IpAdress)

	if err != nil || !result.Result {
		return result, err
	}

	return result, err
}

func PrepareNewsLetterSubscription(nr *newsLetterRepository, c context.Context, filter domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	for _, ID := range filter.StoreID {

		storeID, err := bson.ObjectIDFromHex(ID)
		if err != nil {
			result.Result = false
			result.Message = "Invalid Store ID"
			return result, err
		}

		lang, err := GetLangugaByCode(c, filter.Lang, nr.database.Collection(localization.CollectionLanguage))
		if err != nil {
			result.Result = false
			result.Message = "Invalid Language Code"
			return result, err
		}

		fields, err := GetFieldsByID(c, storeID, nr.database.Collection(stores.CollectionStore))
		if err != nil {
			result.Result = false
			result.Message = "Failed to get store fields"
			return result, err
		}

		store := (*fields)["_id"]
		suscriptionNews := message.NewsLetterSubscription{
			Guid:         uuid.New().String(),
			Email:        filter.Email,
			Active:       false,
			StoreID:      store.(bson.ObjectID),
			CreatedOnUtc: time.Now(),
			IpAddress:    IpAdress,
			LanguageID:   lang.ID,
		}

		state, err := AddNewsLetterSubscription(nr, c, suscriptionNews)
		if err != nil {
			result.Result = false
			result.Message = "Failed to add subscription"
			return result, err
		}

		// Send email to the user for confirmation to activate the subscription
		if state {

			template, err := GetMessageTemplateEmail(nr, c, "NEWSLETTER_SUBSCRIPTION_ACTIVATION_MESSAGE", lang.ID, nr.database.Collection(message.CollectionMessageTemplate))
			if err != nil {
				return result, err
			}

			tokensSubject := GetTokens(c, template.Subject)
			template.Subject = ReplaceTokens(c, template.Subject, tokensSubject, fields, stores.CollectionStore)

			slugs, err := GetSlugsNewsLetterbyLang(nr, c, message.CollectionNewsLetterSubscription, lang.ID)
			if err != nil {
				return result, err
			}

			tokensBody := GetTokens(c, template.Body)
			template.Body = ReplaceTokens(c, template.Body, tokensBody, fields, stores.CollectionStore)
			template.Body = ReplaceTokens(c, template.Body, tokensBody, slugs, seo.CollectionUrlRecord)

			// Convert suscriptionNews (struct) to bson.M
			doc, err := bson.Marshal(suscriptionNews)
			if err != nil {
				result.Result = false
				result.Message = "Failed to convert suscriptionNews (struct) to bson.M"
				return result, err
			}
			// Convert doc to bson.M
			var bsonMap bson.M
			err = bson.Unmarshal(doc, &bsonMap)
			if err != nil {
				result.Result = false
				result.Message = "Failed to convert suscriptionNews doc to bson.M"
				return result, err
			}

			template.Body = ReplaceTokens(c, template.Body, tokensBody, &bsonMap, message.CollectionNewsLetterSubscription)

			email, err := GetEmailAcount(c, nr.database.Collection(message.CollectionEmailAccount))
			if err != nil || email == nil {
				result.Result = false
				result.Message = "Failed to get email account"
				return result, err
			}

			ok, err := AddQueuedEmail(c, *email, filter.Email, "", "", message.High, template.Subject, template.Body, bson.ObjectID{}, nr.database.Collection(message.CollectionQueuedEmail))
			if err != nil || !ok {
				result.Result = false
				result.Message = "Failed to add queued email"
				return result, err
			}
		}

		locale, _ := GetLocalebyName(c, "Newsletter.SubscribeEmailSent", lang.ID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		result.Result = true
		result.Message = locale.ResourceValue

	}

	return result, err
}

func AddNewsLetterSubscription(nr *newsLetterRepository, c context.Context, suscriptionNews message.NewsLetterSubscription) (bool, error) {
	var suscriber []message.NewsLetterSubscription

	findOptions := options.Find().
		SetLimit(1)

	query := bson.M{
		"email":    suscriptionNews.Email,
		"store_id": suscriptionNews.StoreID,
	}

	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return false, err
	}

	err = cursor.All(c, &suscriber)
	if err != nil {
		return false, err
	}
	defer cursor.Close(c)

	if len(suscriber) > 0 {
		// Check if the subscription is already active
		if suscriber[0].Active {
			locale, err := GetLocalebyName(c, "Newsletter.AlreadySubscribed", suscriber[0].LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
			if err != nil {
				return false, err
			}
			return false, fmt.Errorf("%s", locale.ResourceValue)
		}
	} else {
		// Insert a new subscription
		_, err = collection.InsertOne(c, suscriptionNews)
		if err != nil {
			return false, err
		}
	}

	return true, err
}

func AddQueuedEmail(c context.Context,
	email message.EmailAccount,
	to string,
	cc string,
	bcc string,
	priority int,
	subject string,
	body string,
	attachedID bson.ObjectID,
	collection mongo.Collection) (bool, error) {

	var item message.QueuedEmail

	if priority != message.High {
		priority = int(message.Low)
	}

	item = message.QueuedEmail{
		PriorityID:            priority,
		From:                  email.Email,
		FromName:              email.DisplayName,
		To:                    to,
		ToName:                "",
		ReplyTo:               "",
		ReplyToName:           "",
		CC:                    cc,
		Bcc:                   bcc,
		Subject:               subject,
		Body:                  body,
		AttachmentFilePath:    "",
		AttachmentFileName:    "",
		AttachedDownloadID:    attachedID,
		CreatedOnUtc:          time.Now(),
		DontSendBeforeDateUtc: nil,
		SentTries:             3,
		SentOnUtc:             nil,
		EmailAccountID:        email.ID,
	}

	_, err := collection.InsertOne(c, item)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetEmailAcount(c context.Context, collection mongo.Collection) (*message.EmailAccount, error) {
	email := []message.EmailAccount{}
	err := error(nil)

	findOptions := options.Find().
		SetLimit(1)

	query := bson.M{}
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &email)
	if err != nil {
		return nil, err
	}

	return &email[0], err
}

func GetMessageTemplateEmail(nr *newsLetterRepository, c context.Context, name string, landID bson.ObjectID, collection mongo.Collection) (message.MessageTemplate, error) {
	err := error(nil)

	query := bson.M{
		"name": name,
	}

	template := message.MessageTemplate{}
	err = collection.FindOne(c, query).Decode(&template)
	if err != nil {
		return template, err
	}

	record, err := GetRercordBySystemName(c, message.CollectionMessageTemplate, nr.database.Collection(security.CollectionPermissionRecord))
	if err != nil {
		return template, err

	}

	items, err := GetLocalizedProperty(c, record.ID, landID, template.ID.Hex(), nr.database.Collection(localization.CollectionLocalizedProperty))
	if err != nil {
		return template, err
	}

	for i := range items {
		switch items[i].LocaleKey {
		case "subject":
			template.Subject = items[i].LocaleValue
		case "body":
			template.Body = items[i].LocaleValue
		}
	}

	return template, err
}

func GetTokensTemplateEmail(c context.Context, name string, collection mongo.Collection) ([]string, error) {
	err := error(nil)

	query := bson.M{
		"name": name,
	}

	template := message.MessageTemplate{}
	err = collection.FindOne(c, query).Decode(&template)
	if err != nil {
		return nil, err
	}

	joined := append(GetTokens(c, template.Subject), GetTokens(c, template.Body)...)

	return joined, err
}

func GetTokens(c context.Context, text string) []string {
	var items []string

	// Regex to find all words between %%
	re := regexp.MustCompile(`%([^%]+)%`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) > 1 {
			items = append(items, "%"+match[1]+"%")
		}
	}

	return items
}

func ReplaceTokens(c context.Context, text string, key []string, values *bson.M, subfix string) string {
	var item string
	item = text
	subfixReplace := subfix + "."
	// Iterate over the keys and replace them with their corresponding values
	for _, k := range key {
		kk := k[1 : len(k)-1]

		if subfix == strings.Split(kk, ".")[0] {
			// Remove the % from the key
			kk = kk[len(subfixReplace):]

			// Check if the key exists in the map and replace it with the value
			if value, ok := (*values)[kk]; ok {
				// Convert the value to a string and replace the token in the text
				valueStr := ToStringAlways(value)
				item = strings.Replace(item, k, valueStr, -1)
			}
		}
	}

	return item
}

func (cu *newsLetterRepository) GetSlugs(c context.Context, record string) ([]string, error) {

	urlRecord, err := GetRercordBySystemName(c, record, cu.database.Collection(security.CollectionPermissionRecord))
	if err != nil {
		return nil, err
	}

	slugs := []string{}
	urls, err := GetSlugsByRecord(c, urlRecord.ID, cu.database.Collection(seo.CollectionUrlRecord))
	if err != nil {
		return nil, err
	}

	for _, url := range urls {
		slugs = append(slugs, url.Slug)
	}

	return slugs, nil
}

func GetSlugsNewsLetterbyLang(cu *newsLetterRepository, c context.Context, record string, lang bson.ObjectID) (*bson.M, error) {

	urlRecord, err := GetRercordBySystemName(c, record, cu.database.Collection(security.CollectionPermissionRecord))
	if err != nil {
		return nil, err
	}

	slugs := bson.M{}
	urls, err := GetSlugsByRecordLang(c, urlRecord.ID, lang, cu.database.Collection(seo.CollectionUrlRecord))
	if err != nil || len(urls) == 0 {
		urls, err = GetSlugsByPermission(c, urlRecord.ID, cu.database.Collection(seo.CollectionUrlRecord))
		if err != nil {
			return nil, err
		}
	}

	if len(urls) > 0 {
		slugs = urls[0]
	}

	return &slugs, err
}
