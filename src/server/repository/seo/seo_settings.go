package repository

import (
	"context"

	domain "earnforglance/server/domain/seo"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type seosettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewSeoSettingsRepository(db mongo.Database, collection string) domain.SeoSettingsRepository {
	return &seosettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *seosettingsRepository) CreateMany(c context.Context, items []domain.SeoSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *seosettingsRepository) Create(c context.Context, seosettings *domain.SeoSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, seosettings)

	return err
}

func (ur *seosettingsRepository) Update(c context.Context, seosettings *domain.SeoSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": seosettings.ID}
	update := bson.M{
		"$set": seosettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *seosettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *seosettingsRepository) Fetch(c context.Context) ([]domain.SeoSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var seosettingss []domain.SeoSettings

	err = cursor.All(c, &seosettingss)
	if seosettingss == nil {
		return []domain.SeoSettings{}, err
	}

	return seosettingss, err
}

func (tr *seosettingsRepository) FetchByID(c context.Context, seosettingsID string) (domain.SeoSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var seosettings domain.SeoSettings

	idHex, err := primitive.ObjectIDFromHex(seosettingsID)
	if err != nil {
		return seosettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&seosettings)
	return seosettings, err
}
