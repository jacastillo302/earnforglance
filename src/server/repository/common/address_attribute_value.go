package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *addressattributevalueRepository) CreateMany(c context.Context, items []domain.AddressAttributeValue) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
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

func (ur *addressattributevalueRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *addressattributevalueRepository) Fetch(c context.Context) ([]domain.AddressAttributeValue, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(addressattributevalueID)
	if err != nil {
		return addressattributevalue, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&addressattributevalue)
	return addressattributevalue, err
}
