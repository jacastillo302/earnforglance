package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customersettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerSettingsRepository(db mongo.Database, collection string) domain.CustomerSettingsRepository {
	return &customersettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customersettingsRepository) CreateMany(c context.Context, items []domain.CustomerSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *customersettingsRepository) Create(c context.Context, customersetting *domain.CustomerSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customersetting)

	return err
}

func (ur *customersettingsRepository) Update(c context.Context, customersetting *domain.CustomerSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customersetting.ID}
	update := bson.M{
		"$set": customersetting,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customersettingsRepository) Delete(c context.Context, customersetting string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customersetting}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *customersettingsRepository) Fetch(c context.Context) ([]domain.CustomerSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customersettings []domain.CustomerSettings

	err = cursor.All(c, &customersettings)
	if customersettings == nil {
		return []domain.CustomerSettings{}, err
	}

	return customersettings, err
}

func (tr *customersettingsRepository) FetchByID(c context.Context, customersettingID string) (domain.CustomerSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var customersetting domain.CustomerSettings

	idHex, err := primitive.ObjectIDFromHex(customersettingID)
	if err != nil {
		return customersetting, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customersetting)
	return customersetting, err
}
