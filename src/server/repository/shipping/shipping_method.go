package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shippingmethodRepository struct {
	database   mongo.Database
	collection string
}

func NewShippingMethodRepository(db mongo.Database, collection string) domain.ShippingMethodRepository {
	return &shippingmethodRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shippingmethodRepository) Create(c context.Context, shippingmethod *domain.ShippingMethod) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shippingmethod)

	return err
}

func (ur *shippingmethodRepository) Update(c context.Context, shippingmethod *domain.ShippingMethod) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shippingmethod.ID}
	update := bson.M{
		"$set": shippingmethod,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shippingmethodRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shippingmethodRepository) Fetch(c context.Context) ([]domain.ShippingMethod, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shippingmethods []domain.ShippingMethod

	err = cursor.All(c, &shippingmethods)
	if shippingmethods == nil {
		return []domain.ShippingMethod{}, err
	}

	return shippingmethods, err
}

func (tr *shippingmethodRepository) FetchByID(c context.Context, shippingmethodID string) (domain.ShippingMethod, error) {
	collection := tr.database.Collection(tr.collection)

	var shippingmethod domain.ShippingMethod

	idHex, err := primitive.ObjectIDFromHex(shippingmethodID)
	if err != nil {
		return shippingmethod, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shippingmethod)
	return shippingmethod, err
}
