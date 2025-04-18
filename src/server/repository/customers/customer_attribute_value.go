package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerAttributeValueRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerAttributeValueRepository(db mongo.Database, collection string) domain.CustomerAttributeValueRepository {
	return &CustomerAttributeValueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *CustomerAttributeValueRepository) CreateMany(c context.Context, items []domain.CustomerAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *CustomerAttributeValueRepository) Create(c context.Context, CustomerAttributeValue *domain.CustomerAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, CustomerAttributeValue)

	return err
}

func (ur *CustomerAttributeValueRepository) Update(c context.Context, CustomerAttributeValue *domain.CustomerAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": CustomerAttributeValue.ID}
	update := bson.M{
		"$set": CustomerAttributeValue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *CustomerAttributeValueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *CustomerAttributeValueRepository) Fetch(c context.Context) ([]domain.CustomerAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var CustomerAttributeValues []domain.CustomerAttributeValue

	err = cursor.All(c, &CustomerAttributeValues)
	if CustomerAttributeValues == nil {
		return []domain.CustomerAttributeValue{}, err
	}

	return CustomerAttributeValues, err
}

func (tr *CustomerAttributeValueRepository) FetchByID(c context.Context, CustomerAttributeValueID string) (domain.CustomerAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var CustomerAttributeValue domain.CustomerAttributeValue

	idHex, err := primitive.ObjectIDFromHex(CustomerAttributeValueID)
	if err != nil {
		return CustomerAttributeValue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&CustomerAttributeValue)
	return CustomerAttributeValue, err
}
