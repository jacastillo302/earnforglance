package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shippingmethodcountrymappingRepository struct {
	database   mongo.Database
	collection string
}

func NewShippingMethodCountryMappingRepository(db mongo.Database, collection string) domain.ShippingMethodCountryMappingRepository {
	return &shippingmethodcountrymappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shippingmethodcountrymappingRepository) Create(c context.Context, shippingmethodcountrymapping *domain.ShippingMethodCountryMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shippingmethodcountrymapping)

	return err
}

func (ur *shippingmethodcountrymappingRepository) Update(c context.Context, shippingmethodcountrymapping *domain.ShippingMethodCountryMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shippingmethodcountrymapping.ID}
	update := bson.M{
		"$set": shippingmethodcountrymapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shippingmethodcountrymappingRepository) Delete(c context.Context, shippingmethodcountrymapping *domain.ShippingMethodCountryMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shippingmethodcountrymapping.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *shippingmethodcountrymappingRepository) Fetch(c context.Context) ([]domain.ShippingMethodCountryMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shippingmethodcountrymappings []domain.ShippingMethodCountryMapping

	err = cursor.All(c, &shippingmethodcountrymappings)
	if shippingmethodcountrymappings == nil {
		return []domain.ShippingMethodCountryMapping{}, err
	}

	return shippingmethodcountrymappings, err
}

func (tr *shippingmethodcountrymappingRepository) FetchByID(c context.Context, shippingmethodcountrymappingID string) (domain.ShippingMethodCountryMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var shippingmethodcountrymapping domain.ShippingMethodCountryMapping

	idHex, err := primitive.ObjectIDFromHex(shippingmethodcountrymappingID)
	if err != nil {
		return shippingmethodcountrymapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shippingmethodcountrymapping)
	return shippingmethodcountrymapping, err
}
