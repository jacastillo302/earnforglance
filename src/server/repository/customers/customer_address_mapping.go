package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customeraddressmappingRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerAddressMappingRepository(db mongo.Database, collection string) domain.CustomerAddressMappingRepository {
	return &customeraddressmappingRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customeraddressmappingRepository) Create(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customeraddressmapping)

	return err
}

func (ur *customeraddressmappingRepository) Update(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customeraddressmapping.ID}
	update := bson.M{
		"$set": customeraddressmapping,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customeraddressmappingRepository) Delete(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customeraddressmapping.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *customeraddressmappingRepository) Fetch(c context.Context) ([]domain.CustomerAddressMapping, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customeraddressmappings []domain.CustomerAddressMapping

	err = cursor.All(c, &customeraddressmappings)
	if customeraddressmappings == nil {
		return []domain.CustomerAddressMapping{}, err
	}

	return customeraddressmappings, err
}

func (tr *customeraddressmappingRepository) FetchByID(c context.Context, customeraddressmappingID string) (domain.CustomerAddressMapping, error) {
	collection := tr.database.Collection(tr.collection)

	var customeraddressmapping domain.CustomerAddressMapping

	idHex, err := primitive.ObjectIDFromHex(customeraddressmappingID)
	if err != nil {
		return customeraddressmapping, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customeraddressmapping)
	return customeraddressmapping, err
}
