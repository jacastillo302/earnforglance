package repository

import (
	"context"

	domain "earnforglance/server/domain/localization"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type localizationsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewLocalizationSettingsRepository(db mongo.Database, collection string) domain.LocalizationSettingsRepository {
	return &localizationsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *localizationsettingsRepository) Create(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, localizationsettings)

	return err
}

func (ur *localizationsettingsRepository) Update(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": localizationsettings.ID}
	update := bson.M{
		"$set": localizationsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *localizationsettingsRepository) Delete(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": localizationsettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *localizationsettingsRepository) Fetch(c context.Context) ([]domain.LocalizationSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var localizationsettingss []domain.LocalizationSettings

	err = cursor.All(c, &localizationsettingss)
	if localizationsettingss == nil {
		return []domain.LocalizationSettings{}, err
	}

	return localizationsettingss, err
}

func (tr *localizationsettingsRepository) FetchByID(c context.Context, localizationsettingsID string) (domain.LocalizationSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var localizationsettings domain.LocalizationSettings

	idHex, err := primitive.ObjectIDFromHex(localizationsettingsID)
	if err != nil {
		return localizationsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&localizationsettings)
	return localizationsettings, err
}
