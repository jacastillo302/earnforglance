package repository

import (
	"context"
	security "earnforglance/server/domain/security"
	domain "earnforglance/server/domain/seo"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetRercordByID(c context.Context, ID string, collection mongo.Collection) (domain.UrlRecord, error) {
	var urlrecord domain.UrlRecord

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return urlrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&urlrecord)
	return urlrecord, err
}

func GetRercordBySystemName(c context.Context, name string, collection mongo.Collection) (security.PermissionRecord, error) {
	var urlrecord security.PermissionRecord
	err := collection.FindOne(c, bson.M{"system_name": name}).Decode(&urlrecord)
	return urlrecord, err
}
