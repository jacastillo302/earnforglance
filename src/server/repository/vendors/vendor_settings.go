package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type vendorsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewVendorSettingsRepository(db mongo.Database, collection string) domain.VendorSettingsRepository {
	return &vendorsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *vendorsettingsRepository) CreateMany(c context.Context, items []domain.VendorSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *vendorsettingsRepository) Create(c context.Context, vendorsettings *domain.VendorSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, vendorsettings)

	return err
}

func (ur *vendorsettingsRepository) Update(c context.Context, vendorsettings *domain.VendorSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendorsettings.ID}
	update := bson.M{
		"$set": vendorsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *vendorsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *vendorsettingsRepository) Fetch(c context.Context) ([]domain.VendorSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var vendorsettingss []domain.VendorSettings

	err = cursor.All(c, &vendorsettingss)
	if vendorsettingss == nil {
		return []domain.VendorSettings{}, err
	}

	return vendorsettingss, err
}

func (tr *vendorsettingsRepository) FetchByID(c context.Context, vendorsettingsID string) (domain.VendorSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var vendorsettings domain.VendorSettings

	idHex, err := bson.ObjectIDFromHex(vendorsettingsID)
	if err != nil {
		return vendorsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendorsettings)
	return vendorsettings, err
}
