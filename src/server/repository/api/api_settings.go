package repository

import (
	"context"

	domain "earnforglance/server/domain/api"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type apiSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewApiSettingsRepository(db mongo.Database, collection string) domain.ApiSettingsRepository {
	return &apiSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *apiSettingsRepository) Create(c context.Context, apiSettings *domain.ApiSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, apiSettings)

	return err
}

func (ur *apiSettingsRepository) CreateMany(c context.Context, items []domain.ApiSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *apiSettingsRepository) Update(c context.Context, apiSettings *domain.ApiSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": apiSettings.ID}
	update := bson.M{
		"$set": apiSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *apiSettingsRepository) Delete(c context.Context, apiSettings string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": apiSettings}
	_, err := collection.DeleteOne(c, filter)
	return err

}

func (ur *apiSettingsRepository) Fetch(c context.Context) ([]domain.ApiSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var apiSettingsList []domain.ApiSettings

	err = cursor.All(c, &apiSettingsList)
	if apiSettingsList == nil {
		return []domain.ApiSettings{}, err
	}

	return apiSettingsList, err
}

func (tr *apiSettingsRepository) FetchByID(c context.Context, apiSettingsID string) (domain.ApiSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var apiSettings domain.ApiSettings

	idHex, err := primitive.ObjectIDFromHex(apiSettingsID)
	if err != nil {
		return apiSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&apiSettings)
	return apiSettings, err
}
