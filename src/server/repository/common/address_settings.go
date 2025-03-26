package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type addresssettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewAddressSettingsRepository(db mongo.Database, collection string) domain.AddressSettingsRepository {
	return &addresssettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *addresssettingsRepository) CreateMany(c context.Context, items []domain.AddressSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *addresssettingsRepository) Create(c context.Context, addresssettings *domain.AddressSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, addresssettings)

	return err
}

func (ur *addresssettingsRepository) Update(c context.Context, addresssettings *domain.AddressSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": addresssettings.ID}
	update := bson.M{
		"$set": addresssettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *addresssettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *addresssettingsRepository) Fetch(c context.Context) ([]domain.AddressSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var addresssettingss []domain.AddressSettings

	err = cursor.All(c, &addresssettingss)
	if addresssettingss == nil {
		return []domain.AddressSettings{}, err
	}

	return addresssettingss, err
}

func (tr *addresssettingsRepository) FetchByID(c context.Context, addresssettingsID string) (domain.AddressSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var addresssettings domain.AddressSettings

	idHex, err := primitive.ObjectIDFromHex(addresssettingsID)
	if err != nil {
		return addresssettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&addresssettings)
	return addresssettings, err
}
