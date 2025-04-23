package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type externalauthenticationsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewExternalAuthenticationSettingsRepository(db mongo.Database, collection string) domain.ExternalAuthenticationSettingsRepository {
	return &externalauthenticationsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *externalauthenticationsettingsRepository) CreateMany(c context.Context, items []domain.ExternalAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *externalauthenticationsettingsRepository) Create(c context.Context, externalauthenticationsettings *domain.ExternalAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, externalauthenticationsettings)

	return err
}

func (ur *externalauthenticationsettingsRepository) Update(c context.Context, externalauthenticationsettings *domain.ExternalAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": externalauthenticationsettings.ActiveAuthenticationMethodSystemNames}
	update := bson.M{
		"$set": externalauthenticationsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *externalauthenticationsettingsRepository) Delete(c context.Context, externalauthenticationsettings string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": externalauthenticationsettings}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *externalauthenticationsettingsRepository) Fetch(c context.Context) ([]domain.ExternalAuthenticationSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var externalauthenticationsettingss []domain.ExternalAuthenticationSettings

	err = cursor.All(c, &externalauthenticationsettingss)
	if externalauthenticationsettingss == nil {
		return []domain.ExternalAuthenticationSettings{}, err
	}

	return externalauthenticationsettingss, err
}

func (tr *externalauthenticationsettingsRepository) FetchByID(c context.Context, externalauthenticationsettingsID string) (domain.ExternalAuthenticationSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var externalauthenticationsettings domain.ExternalAuthenticationSettings

	idHex, err := bson.ObjectIDFromHex(externalauthenticationsettingsID)
	if err != nil {
		return externalauthenticationsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&externalauthenticationsettings)
	return externalauthenticationsettings, err
}
