package repository

import (
	"context"

	domain "earnforglance/server/domain/stores"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storeRepository struct {
	database   mongo.Database
	collection string
}

func NewStoreRepository(db mongo.Database, collection string) domain.StoreRepository {
	return &storeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *storeRepository) CreateMany(c context.Context, items []domain.Store) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *storeRepository) Create(c context.Context, store *domain.Store) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, store)

	return err
}

func (ur *storeRepository) Update(c context.Context, store *domain.Store) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": store.ID}
	update := bson.M{
		"$set": store,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *storeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *storeRepository) Fetch(c context.Context) ([]domain.Store, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var stores []domain.Store

	err = cursor.All(c, &stores)
	if stores == nil {
		return []domain.Store{}, err
	}

	return stores, err
}

func (tr *storeRepository) FetchByID(c context.Context, storeID string) (domain.Store, error) {
	collection := tr.database.Collection(tr.collection)

	var store domain.Store

	idHex, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return store, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&store)
	return store, err
}
