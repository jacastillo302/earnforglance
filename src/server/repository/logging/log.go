package repository

import (
	"context"

	domain "earnforglance/server/domain/logging"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logRepository struct {
	database   mongo.Database
	collection string
}

func NewLogRepository(db mongo.Database, collection string) domain.LogRepository {
	return &logRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *logRepository) CreateMany(c context.Context, items []domain.Log) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *logRepository) Create(c context.Context, log *domain.Log) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, log)

	return err
}

func (ur *logRepository) Update(c context.Context, log *domain.Log) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": log.ID}
	update := bson.M{
		"$set": log,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *logRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *logRepository) Fetch(c context.Context) ([]domain.Log, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var logs []domain.Log

	err = cursor.All(c, &logs)
	if logs == nil {
		return []domain.Log{}, err
	}

	return logs, err
}

func (tr *logRepository) FetchByID(c context.Context, logID string) (domain.Log, error) {
	collection := tr.database.Collection(tr.collection)

	var log domain.Log

	idHex, err := primitive.ObjectIDFromHex(logID)
	if err != nil {
		return log, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&log)
	return log, err
}
