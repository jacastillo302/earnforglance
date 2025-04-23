package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type downloadRepository struct {
	database   mongo.Database
	collection string
}

func NewDownloadRepository(db mongo.Database, collection string) domain.DownloadRepository {
	return &downloadRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *downloadRepository) CreateMany(c context.Context, items []domain.Download) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *downloadRepository) Create(c context.Context, download *domain.Download) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, download)

	return err
}

func (ur *downloadRepository) Update(c context.Context, download *domain.Download) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": download.ID}
	update := bson.M{
		"$set": download,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *downloadRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *downloadRepository) Fetch(c context.Context) ([]domain.Download, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var downloads []domain.Download

	err = cursor.All(c, &downloads)
	if downloads == nil {
		return []domain.Download{}, err
	}

	return downloads, err
}

func (tr *downloadRepository) FetchByID(c context.Context, downloadID string) (domain.Download, error) {
	collection := tr.database.Collection(tr.collection)

	var download domain.Download

	idHex, err := bson.ObjectIDFromHex(downloadID)
	if err != nil {
		return download, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&download)
	return download, err
}
