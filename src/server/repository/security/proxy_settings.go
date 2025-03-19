package repository

import (
	"context"

	domain "earnforglance/server/domain/security"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type proxysettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewProxySettingsRepository(db mongo.Database, collection string) domain.ProxySettingsRepository {
	return &proxysettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *proxysettingsRepository) Create(c context.Context, proxysettings *domain.ProxySettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, proxysettings)

	return err
}

func (ur *proxysettingsRepository) Update(c context.Context, proxysettings *domain.ProxySettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": proxysettings.ID}
	update := bson.M{
		"$set": proxysettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *proxysettingsRepository) Delete(c context.Context, proxysettings *domain.ProxySettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": proxysettings.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *proxysettingsRepository) Fetch(c context.Context) ([]domain.ProxySettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var proxysettingss []domain.ProxySettings

	err = cursor.All(c, &proxysettingss)
	if proxysettingss == nil {
		return []domain.ProxySettings{}, err
	}

	return proxysettingss, err
}

func (tr *proxysettingsRepository) FetchByID(c context.Context, proxysettingsID string) (domain.ProxySettings, error) {
	collection := tr.database.Collection(tr.collection)

	var proxysettings domain.ProxySettings

	idHex, err := primitive.ObjectIDFromHex(proxysettingsID)
	if err != nil {
		return proxysettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&proxysettings)
	return proxysettings, err
}
