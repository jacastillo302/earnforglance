package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type messagesSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMessagesSettingsRepository(db mongo.Database, collection string) domain.MessagesSettingsRepository {
	return &messagesSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *messagesSettingsRepository) CreateMany(c context.Context, items []domain.MessagesSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *messagesSettingsRepository) Create(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessagesSettings)

	return err
}

func (ur *messagesSettingsRepository) Update(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessagesSettings.ID}
	update := bson.M{
		"$set": MessagesSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *messagesSettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *messagesSettingsRepository) Fetch(c context.Context) ([]domain.MessagesSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var MessagesSettingss []domain.MessagesSettings

	err = cursor.All(c, &MessagesSettingss)
	if MessagesSettingss == nil {
		return []domain.MessagesSettings{}, err
	}

	return MessagesSettingss, err
}

func (tr *messagesSettingsRepository) FetchByID(c context.Context, MessagesSettingsID string) (domain.MessagesSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var MessagesSettings domain.MessagesSettings

	idHex, err := bson.ObjectIDFromHex(MessagesSettingsID)
	if err != nil {
		return MessagesSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessagesSettings)
	return MessagesSettings, err
}
