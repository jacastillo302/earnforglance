package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type producteditorsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewProductEditorSettingsRepository(db mongo.Database, collection string) domain.ProductEditorSettingsRepository {
	return &producteditorsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *producteditorsettingsRepository) CreateMany(c context.Context, items []domain.ProductEditorSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *producteditorsettingsRepository) Create(c context.Context, producteditorsettings *domain.ProductEditorSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, producteditorsettings)

	return err
}

func (ur *producteditorsettingsRepository) Update(c context.Context, producteditorsettings *domain.ProductEditorSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": producteditorsettings.ID}
	update := bson.M{
		"$set": producteditorsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *producteditorsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *producteditorsettingsRepository) Fetch(c context.Context) ([]domain.ProductEditorSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var producteditorsettingss []domain.ProductEditorSettings

	err = cursor.All(c, &producteditorsettingss)
	if producteditorsettingss == nil {
		return []domain.ProductEditorSettings{}, err
	}

	return producteditorsettingss, err
}

func (tr *producteditorsettingsRepository) FetchByID(c context.Context, producteditorsettingsID string) (domain.ProductEditorSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var producteditorsettings domain.ProductEditorSettings

	idHex, err := bson.ObjectIDFromHex(producteditorsettingsID)
	if err != nil {
		return producteditorsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&producteditorsettings)
	return producteditorsettings, err
}
