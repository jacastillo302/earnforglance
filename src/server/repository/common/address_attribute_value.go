package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type addressattributevalueRepository struct {
	database   mongo.Database
	collection string
}

func NewAddressAttributeValueRepository(db mongo.Database, collection string) domain.AddressAttributeValueRepository {
	return &addressattributevalueRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *addressattributevalueRepository) Create(c context.Context, addressattributevalue *domain.AddressAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, addressattributevalue)

	return err
}

func (ur *addressattributevalueRepository) Update(c context.Context, addressattributevalue *domain.AddressAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": addressattributevalue.ID}
	update := bson.M{
		"$set": addressattributevalue,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *addressattributevalueRepository) Delete(c context.Context, addressattributevalue *domain.AddressAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": addressattributevalue.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *addressattributevalueRepository) Fetch(c context.Context) ([]domain.AddressAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var addressattributevalues []domain.AddressAttributeValue

	err = cursor.All(c, &addressattributevalues)
	if addressattributevalues == nil {
		return []domain.AddressAttributeValue{}, err
	}

	return addressattributevalues, err
}

func (tr *addressattributevalueRepository) FetchByID(c context.Context, addressattributevalueID string) (domain.AddressAttributeValue, error) {
	collection := tr.database.Collection(tr.collection)

	var addressattributevalue domain.AddressAttributeValue

	idHex, err := primitive.ObjectIDFromHex(addressattributevalueID)
	if err != nil {
		return addressattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&addressattributevalue)
	return addressattributevalue, err
}
