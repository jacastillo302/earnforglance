package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sitemapsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewSitemapSettingsRepository(db mongo.Database, collection string) domain.SitemapSettingsRepository {
	return &sitemapsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *sitemapsettingsRepository) CreateMany(c context.Context, items []domain.SitemapSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *sitemapsettingsRepository) Create(c context.Context, sitemapsettings *domain.SitemapSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, sitemapsettings)

	return err
}

func (ur *sitemapsettingsRepository) Update(c context.Context, sitemapsettings *domain.SitemapSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": sitemapsettings.ID}
	update := bson.M{
		"$set": sitemapsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *sitemapsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *sitemapsettingsRepository) Fetch(c context.Context) ([]domain.SitemapSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var sitemapsettingss []domain.SitemapSettings

	err = cursor.All(c, &sitemapsettingss)
	if sitemapsettingss == nil {
		return []domain.SitemapSettings{}, err
	}

	return sitemapsettingss, err
}

func (tr *sitemapsettingsRepository) FetchByID(c context.Context, sitemapsettingsID string) (domain.SitemapSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var sitemapsettings domain.SitemapSettings

	idHex, err := primitive.ObjectIDFromHex(sitemapsettingsID)
	if err != nil {
		return sitemapsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&sitemapsettings)
	return sitemapsettings, err
}
