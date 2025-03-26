package repository

import (
	"context"

	domain "earnforglance/server/domain/gdpr"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type gdprsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewGdprSettingsRepository(db mongo.Database, collection string) domain.GdprSettingsRepository {
	return &gdprsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *gdprsettingsRepository) CreateMany(c context.Context, items []domain.GdprSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *gdprsettingsRepository) Create(c context.Context, gdprsettings *domain.GdprSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, gdprsettings)

	return err
}

func (ur *gdprsettingsRepository) Update(c context.Context, gdprsettings *domain.GdprSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": gdprsettings.ID}
	update := bson.M{
		"$set": gdprsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *gdprsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *gdprsettingsRepository) Fetch(c context.Context) ([]domain.GdprSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var gdprsettingss []domain.GdprSettings

	err = cursor.All(c, &gdprsettingss)
	if gdprsettingss == nil {
		return []domain.GdprSettings{}, err
	}

	return gdprsettingss, err
}

func (tr *gdprsettingsRepository) FetchByID(c context.Context, gdprsettingsID string) (domain.GdprSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var gdprsettings domain.GdprSettings

	idHex, err := primitive.ObjectIDFromHex(gdprsettingsID)
	if err != nil {
		return gdprsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&gdprsettings)
	return gdprsettings, err
}
