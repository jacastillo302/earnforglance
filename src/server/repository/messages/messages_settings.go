package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessagesSettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMessagesSettingsRepository(db mongo.Database, collection string) domain.MessagesSettingsRepository {
	return &MessagesSettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *MessagesSettingsRepository) Create(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessagesSettings)

	return err
}

func (ur *MessagesSettingsRepository) Update(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessagesSettings.ID}
	update := bson.M{
		"$set": MessagesSettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *MessagesSettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *MessagesSettingsRepository) Fetch(c context.Context) ([]domain.MessagesSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
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

func (tr *MessagesSettingsRepository) FetchByID(c context.Context, MessagesSettingsID string) (domain.MessagesSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var MessagesSettings domain.MessagesSettings

	idHex, err := primitive.ObjectIDFromHex(MessagesSettingsID)
	if err != nil {
		return MessagesSettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessagesSettings)
	return MessagesSettings, err
}
