package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mediaSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMediaSettingsRepository(db mongo.Database, collection string) domain.MediaSettingsRepository {
	return &mediaSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *mediaSettingsRepository) CreateMany(c context.Context, items []domain.MediaSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *mediaSettingsRepository) Create(c context.Context, MediaSettings *domain.MediaSettings) error {
	_, err := ur.database.Collection(ur.collection).InsertOne(c, MediaSettings)
	return err
}

func (ur *mediaSettingsRepository) Update(c context.Context, MediaSettings *domain.MediaSettings) error {
	filter := bson.M{"_id": MediaSettings.ID}
	update := bson.M{"$set": MediaSettings}
	_, err := ur.database.Collection(ur.collection).UpdateOne(c, filter, update)
	return err
}

func (ur *mediaSettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *mediaSettingsRepository) Fetch(c context.Context) ([]domain.MediaSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var mediaSettings []domain.MediaSettings

	err = cursor.All(c, &mediaSettings)
	if mediaSettings == nil {
		return []domain.MediaSettings{}, err
	}

	return mediaSettings, err
}

func (tr *mediaSettingsRepository) FetchByID(c context.Context, MediaSettingsID string) (domain.MediaSettings, error) {
	var MediaSettings domain.MediaSettings
	idHex, err := primitive.ObjectIDFromHex(MediaSettingsID)
	if err != nil {
		return MediaSettings, err
	}
	err = tr.database.Collection(tr.collection).FindOne(c, bson.M{"_id": idHex}).Decode(&MediaSettings)
	return MediaSettings, err
}
