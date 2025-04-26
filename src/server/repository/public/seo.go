package repository

import (
	"context"
	domain "earnforglance/server/domain/seo"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUrlRercord(c context.Context, urlrecordID string, collection mongo.Collection) (domain.UrlRecord, error) {
	var urlrecord domain.UrlRecord

	idHex, err := bson.ObjectIDFromHex(urlrecordID)
	if err != nil {
		return urlrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&urlrecord)
	return urlrecord, err
}

func GetSlugsByRecord(c context.Context, recordID bson.ObjectID, collection mongo.Collection) ([]domain.UrlRecord, error) {
	var urlrecord []domain.UrlRecord

	query := bson.M{
		"entity_id": recordID,
	}

	cursor, err := collection.Find(c, query)
	if err != nil {
		return urlrecord, err
	}

	err = cursor.All(c, &urlrecord)
	if err != nil {
		return urlrecord, err
	}
	defer cursor.Close(c)

	return urlrecord, nil
}
