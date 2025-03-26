package repository

import (
	"context"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pdfsettingsRepository struct {
	database   mongo.Database
	collection string
}

func NewPdfSettingsRepository(db mongo.Database, collection string) domain.PdfSettingsRepository {
	return &pdfsettingsRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pdfsettingsRepository) CreateMany(c context.Context, items []domain.PdfSettings) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *pdfsettingsRepository) Create(c context.Context, pdfsettings *domain.PdfSettings) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, pdfsettings)

	return err
}

func (ur *pdfsettingsRepository) Update(c context.Context, pdfsettings *domain.PdfSettings) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": pdfsettings.ID}
	update := bson.M{
		"$set": pdfsettings,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *pdfsettingsRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pdfsettingsRepository) Fetch(c context.Context) ([]domain.PdfSettings, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var pdfsettingss []domain.PdfSettings

	err = cursor.All(c, &pdfsettingss)
	if pdfsettingss == nil {
		return []domain.PdfSettings{}, err
	}

	return pdfsettingss, err
}

func (tr *pdfsettingsRepository) FetchByID(c context.Context, pdfsettingsID string) (domain.PdfSettings, error) {
	collection := tr.database.Collection(tr.collection)

	var pdfsettings domain.PdfSettings

	idHex, err := primitive.ObjectIDFromHex(pdfsettingsID)
	if err != nil {
		return pdfsettings, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&pdfsettings)
	return pdfsettings, err
}
