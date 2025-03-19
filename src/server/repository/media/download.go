package repository

import (
	"context"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ur *downloadRepository) Delete(c context.Context, download *domain.Download) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": download.ID}
	_, err := collection.DeleteOne(c, filter)
	return err

}

func (ur *downloadRepository) Fetch(c context.Context) ([]domain.Download, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := primitive.ObjectIDFromHex(downloadID)
	if err != nil {
		return download, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&download)
	return download, err
}
