package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *shippingmethodcountrymappingRepository) CreateMany(c context.Context, items []domain.ShippingMethodCountryMapping) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
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

func (ur *shippingmethodcountrymappingRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shippingmethodcountrymappingRepository) Fetch(c context.Context) ([]domain.ShippingMethodCountryMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(shippingmethodcountrymappingID)
	if err != nil {
		return shippingmethodcountrymapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shippingmethodcountrymapping)
	return shippingmethodcountrymapping, err
}
