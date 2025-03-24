package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sitemapxmlsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewSitemapXmlSettingsRepository(db mongo.Database, collection string) domain.SitemapXmlSettingsRepository {
	return &sitemapxmlsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *sitemapxmlsettingsRepository) CreateMany(c context.Context, items []domain.SitemapXmlSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *sitemapxmlsettingsRepository) Create(c context.Context, sitemapxmlsettings *domain.SitemapXmlSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, sitemapxmlsettings)

	return err
}

func (ur *sitemapxmlsettingsRepository) Update(c context.Context, sitemapxmlsettings *domain.SitemapXmlSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": sitemapxmlsettings.ID}
	update := bson.M{
		"$set": sitemapxmlsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *sitemapxmlsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *sitemapxmlsettingsRepository) Fetch(c context.Context) ([]domain.SitemapXmlSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var sitemapxmlsettingss []domain.SitemapXmlSettings

	err = cursor.All(c, &sitemapxmlsettingss)
	if sitemapxmlsettingss == nil {
		return []domain.SitemapXmlSettings{}, err
	}

	return sitemapxmlsettingss, err
}

func (tr *sitemapxmlsettingsRepository) FetchByID(c context.Context, sitemapxmlsettingsID string) (domain.SitemapXmlSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var sitemapxmlsettings domain.SitemapXmlSettings

	idHex, err := primitive.ObjectIDFromHex(sitemapxmlsettingsID)
	if err != nil {
		return sitemapxmlsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&sitemapxmlsettings)
	return sitemapxmlsettings, err
}
