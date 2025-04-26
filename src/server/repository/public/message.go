package repository

import (
	"context"
	localization "earnforglance/server/domain/localization"
	message "earnforglance/server/domain/messages"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	seo "earnforglance/server/domain/seo"

	"earnforglance/server/service/data/mongo"
	service "earnforglance/server/service/public"
	"fmt"
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

	news := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	collection.FindOne(c, query).Decode(&news)

	if news.Active {

		update, err := collection.UpdateOne(c, query, bson.M{"$set": bson.M{"active": false, "newsletter_subscription_guid": uuid.New()}})
		if err != nil {
			result.Result = false
			result.Message = "Failed to inactivate subscription"
			return result, err
		}

		if update.MatchedCount == 0 {
			result.Result = false
			result.Message = "No matching subscription found"
			return result, err
		}
	}

	langID, err := GetLangugaByCode(c, lang, nr.database.Collection(localization.CollectionLanguage))
	if err != nil {
		result.Result = false
		result.Message = "Invalid Language Code"
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

	uu_id, err := uuid.Parse(guid)
	if err != nil {
		result.Result = false
		result.Message = "Invalid Guid"
		return result, err
	}

	query := bson.M{
		"newsletter_subscription_guid": uu_id,
	}

	news := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	err = collection.FindOne(c, query).Decode(&news)
	if err != nil {
		result.Result = false
		result.Message = err.Error()
		return result, err
	}

	if news.Guid.String() != guid {
		locale, err := GetLocalebyName(c, "Newsletter.ResultActivated.InvalidGuid", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		result.Result = false
		result.Message = locale.ResourceValue
		return result, err
	}

	if news.Active {
		// Deactivate the subscription
		update, err := collection.UpdateOne(c, query, bson.M{"$set": bson.M{"active": false, "newsletter_subscription_guid": uuid.New()}})
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

	uu_id, err := uuid.Parse(guid)
	if err != nil {
		result.Result = false
		result.Message = "Invalid Guid"
		return result, err
	}

	query := bson.M{
		"newsletter_subscription_guid": uu_id,
	}

	news := message.NewsLetterSubscription{}
	collection := nr.database.Collection(message.CollectionNewsLetterSubscription)
	err = collection.FindOne(c, query).Decode(&news)
	if err != nil {
		result.Result = false
		result.Message = err.Error()
		return result, err
	}

	if news.Guid.String() != guid {
		locale, err := GetLocalebyName(c, "Newsletter.ResultActivated.InvalidGuid", news.LanguageID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		result.Result = false
		result.Message = locale.ResourceValue
		return result, err
	}

	if !news.Active {
		// Activate the subscription
		update, err := collection.UpdateOne(c, query, bson.M{"$set": bson.M{"active": true, "newsletter_subscription_guid": uuid.New()}})
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

	for _, store := range filter.StoreID {

		storeID, err := bson.ObjectIDFromHex(store)
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

		suscriptionNews := message.NewsLetterSubscription{
			Guid:         uuid.New(),
			Email:        filter.Email,
			Active:       false,
			StoreID:      storeID,
			CreatedOnUtc: time.Now(),
			IpAddress:    IpAdress,
			LanguageID:   lang.ID,
		}

		locale, _ := GetLocalebyName(c, "Newsletter.SubscribeEmailSent", lang.ID.Hex(), nr.database.Collection(localization.CollectionLocaleStringResource))
		AddNewsLetterSubscription(nr, c, suscriptionNews)

		tokens, err := service.ReadJsonTokens("StoreTokens")
		if err != nil {
			return result, err
		}

		fmt.Println("tokens", tokens)

		tokens, err = service.ReadJsonTokens("SubscriptionTokens")
		if err != nil {
			return result, err
		}

		fmt.Println("tokens", tokens)

		result.Result = true
		result.Message = locale.ResourceValue

	}

	return result, err
}

func AddNewsLetterSubscription(vr *newsLetterRepository, c context.Context, suscriptionNews message.NewsLetterSubscription) (bool, error) {
	var suscriber []message.NewsLetterSubscription
	result := false

	findOptions := options.Find().
		SetLimit(1)

	query := bson.M{
		"email":    suscriptionNews.Email,
		"store_id": suscriptionNews.StoreID,
	}

	collection := vr.database.Collection(message.CollectionNewsLetterSubscription)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(c, &suscriber)
	if err != nil {
		return result, err
	}

	if len(suscriber) > 0 {
		if suscriber[0].Active {
			return true, err
		}
		return false, err
	}

	_, err = collection.InsertOne(c, suscriptionNews)
	if err != nil {
		return false, err
	}

	_, err = GetMessageTemplateEmail(c, "NEWSLETTER_SUBSCRIPTION_ACTIVATION_MESSAGE", map[string]string{"email": suscriptionNews.Email}, vr.database.Collection(message.CollectionQueuedEmail))
	if err != nil {
		return false, err
	}

	return true, err
}

func AddQueuedEmail(c context.Context, item message.QueuedEmail, collection mongo.Collection) (bool, error) {

	_, err := collection.InsertOne(c, item)

	if err != nil {
		return false, err
	}

	return true, err
}

func GetMessageTemplateEmail(c context.Context, name string, tokens map[string]string, collection mongo.Collection) (message.QueuedEmail, error) {
	var item message.QueuedEmail
	err := error(nil)

	query := bson.M{
		"name": name,
	}

	template := message.MessageTemplate{}
	err = collection.FindOne(c, query).Decode(&template)
	if err != nil {
		return item, err
	}

	//filtered := service.FilterTypesByValue(items, product.ProductTypeID)

	for key, value := range tokens {
		item.Body = ReplaceToken(c, template.Body, key, value, collection)
		item.Subject = ReplaceToken(c, template.Subject, key, value, collection)
	}

	return item, err
}

func ReplaceToken(c context.Context, body string, key string, value string, collection mongo.Collection) string {
	var item string

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
