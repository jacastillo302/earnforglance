package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type checkoutattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewCheckoutAttributeValueRepository(db mongo.Database, collection string) domain.CheckoutAttributeValueRepository {
	return &checkoutattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *checkoutattributevalueRepository) CreateMany(c context.Context, items []domain.CheckoutAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *checkoutattributevalueRepository) Create(c context.Context, checkoutattributevalue *domain.CheckoutAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, checkoutattributevalue)

	return err
}

func (ur *checkoutattributevalueRepository) Update(c context.Context, checkoutattributevalue *domain.CheckoutAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": checkoutattributevalue.ID}
	update := bson.M{
		"$set": checkoutattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *checkoutattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *checkoutattributevalueRepository) Fetch(c context.Context) ([]domain.CheckoutAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var checkoutattributevalues []domain.CheckoutAttributeValue

	err = cursor.All(c, &checkoutattributevalues)
	if checkoutattributevalues == nil {
		return []domain.CheckoutAttributeValue{}, err
	}

	return checkoutattributevalues, err
}

func (tr *checkoutattributevalueRepository) FetchByID(c context.Context, checkoutattributevalueID string) (domain.CheckoutAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var checkoutattributevalue domain.CheckoutAttributeValue

	idHex, err := bson.ObjectIDFromHex(checkoutattributevalueID)
	if err != nil {
		return checkoutattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&checkoutattributevalue)
	return checkoutattributevalue, err
}
