package repository

import (
	"context"
	store "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetDefaultStore(c context.Context, collection mongo.Collection) (*store.Store, error) {
	var stores []store.Store
	var store *store.Store
	store = nil

	findOptions := options.Find().
		SetSort(bson.D{{Key: "display_order", Value: 1}}).
		SetLimit(1)

	cursor, err := collection.Find(c, bson.M{"deleted": false}, findOptions)
	if err != nil {
		return store, err
	}

	err = cursor.All(c, &stores)
	if err != nil {
		return store, err
	}
	defer cursor.Close(c)

	if len(stores) > 0 {

		store = &stores[0]
	}

	return store, err
}

func GetStoreByID(c context.Context, ID bson.ObjectID, collection mongo.Collection) (*store.Store, error) {
	var store store.Store
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&store)
	return &store, err
}

func GetFieldsByID(c context.Context, ID bson.ObjectID, collection mongo.Collection) (*bson.M, error) {
	var store bson.M
	err := collection.FindOne(c, bson.M{"_id": ID}).Decode(&store)
	return &store, err
}
