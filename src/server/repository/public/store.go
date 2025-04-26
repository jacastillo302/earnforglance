package repository

import (
	"context"
	store "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetStoreByID(c context.Context, ID bson.ObjectID, collection mongo.Collection) (*store.Store, error) {
	var store store.Store
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&store)
	return &store, err
}
