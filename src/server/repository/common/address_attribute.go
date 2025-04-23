package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type addressattributeRepository struct {
	database   mongo.Database
	collection string
}

func NewAddressAttributeRepository(db mongo.Database, collection string) domain.AddressAttributeRepository {
	return &addressattributeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *addressattributeRepository) CreateMany(c context.Context, items []domain.AddressAttribute) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *addressattributeRepository) Create(c context.Context, addressattribute *domain.AddressAttribute) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, addressattribute)

	return err
}

func (ur *addressattributeRepository) Update(c context.Context, addressattribute *domain.AddressAttribute) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": addressattribute.ID}
	update := bson.M{
		"$set": addressattribute,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *addressattributeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *addressattributeRepository) Fetch(c context.Context) ([]domain.AddressAttribute, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var addressattributes []domain.AddressAttribute

	err = cursor.All(c, &addressattributes)
	if addressattributes == nil {
		return []domain.AddressAttribute{}, err
	}

	return addressattributes, err
}

func (tr *addressattributeRepository) FetchByID(c context.Context, addressattributeID string) (domain.AddressAttribute, error) {
	collection := tr.database.Collection(tr.collection)

	var addressattribute domain.AddressAttribute

	idHex, err := bson.ObjectIDFromHex(addressattributeID)
	if err != nil {
		return addressattribute, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&addressattribute)
	return addressattribute, err
}
