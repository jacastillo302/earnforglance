package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type addressRepository struct {
	database   mongo.Database
	collection string
}

func NewAddressRepository(db mongo.Database, collection string) domain.AddressRepository {
	return &addressRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *addressRepository) CreateMany(c context.Context, items []domain.Address) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *addressRepository) Create(c context.Context, address *domain.Address) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, address)

	return err
}

func (ur *addressRepository) Update(c context.Context, address *domain.Address) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": address.ID}
	update := bson.M{
		"$set": address,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *addressRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *addressRepository) Fetch(c context.Context) ([]domain.Address, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var addresses []domain.Address

	err = cursor.All(c, &addresses)
	if addresses == nil {
		return []domain.Address{}, err
	}

	return addresses, err
}

func (tr *addressRepository) FetchByID(c context.Context, addressID string) (domain.Address, error) {
	collection := tr.database.Collection(tr.collection)

	var address domain.Address

	idHex, err := bson.ObjectIDFromHex(addressID)
	if err != nil {
		return address, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&address)
	return address, err
}
