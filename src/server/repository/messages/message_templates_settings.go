package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type messageTemplatesSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageTemplatesSettingsRepository(db mongo.Database, collection string) domain.MessageTemplatesSettingsRepository {
	return &messageTemplatesSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *messageTemplatesSettingsRepository) CreateMany(c context.Context, items []domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *messageTemplatesSettingsRepository) Create(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessageTemplatesSettings)

	return err
}

func (ur *messageTemplatesSettingsRepository) Update(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplatesSettings.ID}
	update := bson.M{
		"$set": MessageTemplatesSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *messageTemplatesSettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *messageTemplatesSettingsRepository) Fetch(c context.Context) ([]domain.MessageTemplatesSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

func (tr *messageTemplatesSettingsRepository) FetchByID(c context.Context, MessageTemplatesSettingsID string) (domain.MessageTemplatesSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var MessageTemplatesSettings domain.MessageTemplatesSettings

	idHex, err := bson.ObjectIDFromHex(MessageTemplatesSettingsID)
	if err != nil {
		return MessageTemplatesSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessageTemplatesSettings)
	return MessageTemplatesSettings, err
}
