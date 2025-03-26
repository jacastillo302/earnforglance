package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type catalogsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewCatalogSettingsRepository(db mongo.Database, collection string) domain.CatalogSettingsRepository {
	return &catalogsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *catalogsettingsRepository) CreateMany(c context.Context, items []domain.CatalogSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *catalogsettingsRepository) Create(c context.Context, catalogsettings *domain.CatalogSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, catalogsettings)

	return err
}

func (ur *catalogsettingsRepository) Update(c context.Context, catalogsettings *domain.CatalogSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": catalogsettings.ID}
	update := bson.M{
		"$set": catalogsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *catalogsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *catalogsettingsRepository) Fetch(c context.Context) ([]domain.CatalogSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var catalogsettingss []domain.CatalogSettings

	err = cursor.All(c, &catalogsettingss)
	if catalogsettingss == nil {
		return []domain.CatalogSettings{}, err
	}

	return catalogsettingss, err
}

func (tr *catalogsettingsRepository) FetchByID(c context.Context, catalogsettingsID string) (domain.CatalogSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var catalogsettings domain.CatalogSettings

	idHex, err := primitive.ObjectIDFromHex(catalogsettingsID)
	if err != nil {
		return catalogsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&catalogsettings)
	return catalogsettings, err
}
