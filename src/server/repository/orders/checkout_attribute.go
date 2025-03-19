package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type checkoutattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewCheckoutAttributeRepository(db mongo.Database, collection string) domain.CheckoutAttributeRepository {
	return &checkoutattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *checkoutattributeRepository) Create(c context.Context, checkoutattribute *domain.CheckoutAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, checkoutattribute)

	return err
}

func (ur *checkoutattributeRepository) Update(c context.Context, checkoutattribute *domain.CheckoutAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": checkoutattribute.ID}
	update := bson.M{
		"$set": checkoutattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *checkoutattributeRepository) Delete(c context.Context, checkoutattribute *domain.CheckoutAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": checkoutattribute.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *checkoutattributeRepository) Fetch(c context.Context) ([]domain.CheckoutAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var checkoutattributes []domain.CheckoutAttribute

	err = cursor.All(c, &checkoutattributes)
	if checkoutattributes == nil {
		return []domain.CheckoutAttribute{}, err
	}

	return checkoutattributes, err
}

func (tr *checkoutattributeRepository) FetchByID(c context.Context, checkoutattributeID string) (domain.CheckoutAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var checkoutattribute domain.CheckoutAttribute

	idHex, err := primitive.ObjectIDFromHex(checkoutattributeID)
	if err != nil {
		return checkoutattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&checkoutattribute)
	return checkoutattribute, err
}
