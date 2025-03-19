package repository

import (
	"context"

	domain "earnforglance/server/domain/vendors"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ur *vendorsettingsRepository) Delete(c context.Context, vendorsettings *domain.VendorSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": vendorsettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err

}

func (ur *vendorsettingsRepository) Fetch(c context.Context) ([]domain.VendorSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := primitive.ObjectIDFromHex(vendorsettingsID)
	if err != nil {
		return vendorsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&vendorsettings)
	return vendorsettings, err
}
