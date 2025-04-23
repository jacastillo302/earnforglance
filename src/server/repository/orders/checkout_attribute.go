package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *checkoutattributeRepository) CreateMany(c context.Context, items []domain.CheckoutAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
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

func (ur *checkoutattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *checkoutattributeRepository) Fetch(c context.Context) ([]domain.CheckoutAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(checkoutattributeID)
	if err != nil {
		return checkoutattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&checkoutattribute)
	return checkoutattribute, err
}
