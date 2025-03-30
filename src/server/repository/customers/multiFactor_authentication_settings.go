package repository

import (
	"context"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type multifactorauthenticationsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewMultiFactorAuthenticationSettingsRepository(db mongo.Database, collection string) domain.MultiFactorAuthenticationSettingsRepository {
	return &multifactorauthenticationsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *multifactorauthenticationsettingsRepository) CreateMany(c context.Context, items []domain.MultiFactorAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *multifactorauthenticationsettingsRepository) Create(c context.Context, multifactorauthenticationsettings *domain.MultiFactorAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, multifactorauthenticationsettings)

	return err
}

func (ur *multifactorauthenticationsettingsRepository) Update(c context.Context, multifactorauthenticationsettings *domain.MultiFactorAuthenticationSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": multifactorauthenticationsettings.ActiveAuthenticationMethodSystemNames}
	update := bson.M{
		"$set": multifactorauthenticationsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *multifactorauthenticationsettingsRepository) Delete(c context.Context, multifactorauthenticationsettings string) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": multifactorauthenticationsettings}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *multifactorauthenticationsettingsRepository) Fetch(c context.Context) ([]domain.MultiFactorAuthenticationSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var multifactorauthenticationsettingss []domain.MultiFactorAuthenticationSettings

	err = cursor.All(c, &multifactorauthenticationsettingss)
	if multifactorauthenticationsettingss == nil {
		return []domain.MultiFactorAuthenticationSettings{}, err
	}

	return multifactorauthenticationsettingss, err
}

func (tr *multifactorauthenticationsettingsRepository) FetchByID(c context.Context, multifactorauthenticationsettingsID string) (domain.MultiFactorAuthenticationSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var multifactorauthenticationsettings domain.MultiFactorAuthenticationSettings

	idHex, err := primitive.ObjectIDFromHex(multifactorauthenticationsettingsID)
	if err != nil {
		return multifactorauthenticationsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&multifactorauthenticationsettings)
	return multifactorauthenticationsettings, err
}
