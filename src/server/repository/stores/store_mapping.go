package repository

import (
	"context"

	domain "earnforglance/server/domain/stores"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type storemappingRepository struct {
	database   mongo.Database
	collection string
}

func NewStoreMappingRepository(db mongo.Database, collection string) domain.StoreMappingRepository {
	return &storemappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *storemappingRepository) Create(c context.Context, storemapping *domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, storemapping)

	return err
}

func (ur *storemappingRepository) Update(c context.Context, storemapping *domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": storemapping.ID}
	update := bson.M{
		"$set": storemapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *storemappingRepository) Delete(c context.Context, storemapping *domain.StoreMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": storemapping.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *storemappingRepository) Fetch(c context.Context) ([]domain.StoreMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var storemappings []domain.StoreMapping

	err = cursor.All(c, &storemappings)
	if storemappings == nil {
		return []domain.StoreMapping{}, err
	}

	return storemappings, err
}

func (tr *storemappingRepository) FetchByID(c context.Context, storemappingID string) (domain.StoreMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var storemapping domain.StoreMapping

	idHex, err := primitive.ObjectIDFromHex(storemappingID)
	if err != nil {
		return storemapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&storemapping)
	return storemapping, err
}
