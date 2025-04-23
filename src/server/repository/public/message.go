package repository

import (
	"context"
	localization "earnforglance/server/domain/localization"
	message "earnforglance/server/domain/messages"
	domain "earnforglance/server/domain/public"
	"earnforglance/server/service/data/mongo"
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

func (cr *newsLetterRepository) NewsLetterSubscription(c context.Context, filter domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	result, err = PrepareNewsLetterSubscription(cr, c, filter, IpAdress)
	if err != nil || !result.Result {
		return result, err
	}

	result.Result = true
	result.Message = "Success"

	return result, err
}

func PrepareNewsLetterSubscription(nr *newsLetterRepository, c context.Context, filter domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	var result domain.NewsLetterResponse
	err := error(nil)

	storeID, err := bson.ObjectIDFromHex(filter.StoreID)
	if err != nil {
		result.Result = false
		result.Message = "Invalid Store ID"
		return result, err
	}

	lang, err := GetNewsLatterLangugaByCode(nr, c, filter.Lang)
	if err != nil {
		result.Result = false
		result.Message = "Invalid Language Code"
		return result, err
	}

	suscriptionNews := message.NewsLetterSubscription{
		Guid:         uuid.New(),
		Email:        filter.Email,
		Active:       true,
		StoreID:      storeID,
		CreatedOnUtc: time.Now(),
		IpAddress:    IpAdress,
		LanguageID:   lang.ID,
	}

	isNotBanned, err := PreventBannedIP(nr, c, IpAdress)
	if !isNotBanned {
		result.Result = false
		result.Message = "IP Address is banned"
		return result, err
	}

	isSuscribe, err := AddNewsLetterSubscription(nr, c, suscriptionNews)
	if !isSuscribe {
		result.Result = false
		result.Message = "Already subscribed"
		return result, err
	}

	result.Result = true
	result.Message = "Subscription successful"

	return result, err
}

func GetNewsLatterLangugaByCode(vr *newsLetterRepository, c context.Context, lang string) (localization.Language, error) {
	collection := vr.database.Collection(localization.CollectionLanguage)
	var item localization.Language
	err := collection.FindOne(c, bson.M{"unique_seo_code": lang}).Decode(&item)
	return item, err
}

func AddNewsLetterSubscription(vr *newsLetterRepository, c context.Context, suscriptionNews message.NewsLetterSubscription) (bool, error) {
	var suscriber []message.NewsLetterSubscription
	result := false

	findOptions := options.Find().
		SetLimit(1)

	query := bson.M{
		"email": suscriptionNews.Email,
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

		query := bson.M{"_id": suscriber[0].ID}
		var update bson.M

		if suscriber[0].Active {
			return false, err
		} else {
			suscriber[0].Active = true
			update = bson.M{
				"$set": suscriber[0],
			}
			// Update the subscriber to active
			updateResul, err := collection.UpdateOne(c, query, update)
			if err != nil || updateResul.MatchedCount == 0 {
				result = false
				return result, err
			}

			return true, err
		}
	}

	_, err = collection.InsertOne(c, suscriptionNews)
	if err != nil {
		return false, err
	}

	return true, err
}

func PreventBannedIP(vr *newsLetterRepository, c context.Context, IpAddress string) (bool, error) {
	var ipsResult []message.NewsLetterSubscription
	result := true

	// Calculate the time 24 hours ago
	timeLimit := time.Now().Add(-24 * time.Hour)

	findOptions := options.Find().
		SetLimit(1)

	query := bson.M{
		"last_ip_address": IpAddress,
		"created_on_utc":  bson.M{"$gte": timeLimit},
	}

	collection := vr.database.Collection(message.CollectionNewsLetterSubscription)
	cursor, err := collection.Find(c, query, findOptions)
	if err != nil {
		return false, err
	}

	err = cursor.All(c, &ipsResult)
	if err != nil {
		return false, err
	}

	if len(ipsResult) > 0 {
		result = false
	}
	fmt.Println("ipsResult", ipsResult)
	fmt.Println("IpAddress", IpAddress)
	return result, err
}
