package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageTemplatesSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageTemplatesSettingsRepository(db mongo.Database, collection string) domain.MessageTemplatesSettingsRepository {
	return &MessageTemplatesSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *MessageTemplatesSettingsRepository) Create(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessageTemplatesSettings)

	return err
}

func (ur *MessageTemplatesSettingsRepository) Update(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplatesSettings.ID}
	update := bson.M{
		"$set": MessageTemplatesSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *MessageTemplatesSettingsRepository) Delete(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplatesSettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *MessageTemplatesSettingsRepository) Fetch(c context.Context) ([]domain.MessageTemplatesSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var MessageTemplatesSettingss []domain.MessageTemplatesSettings

	err = cursor.All(c, &MessageTemplatesSettingss)
	if MessageTemplatesSettingss == nil {
		return []domain.MessageTemplatesSettings{}, err
	}

	return MessageTemplatesSettingss, err
}

func (tr *MessageTemplatesSettingsRepository) FetchByID(c context.Context, MessageTemplatesSettingsID string) (domain.MessageTemplatesSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var MessageTemplatesSettings domain.MessageTemplatesSettings

	idHex, err := primitive.ObjectIDFromHex(MessageTemplatesSettingsID)
	if err != nil {
		return MessageTemplatesSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessageTemplatesSettings)
	return MessageTemplatesSettings, err
}
