package repository

import (
	"context"

	domain "earnforglance/server/domain/logging"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type activitylogRepository struct {
	database   mongo.Database
	collection string
}

func NewActivityLogRepository(db mongo.Database, collection string) domain.ActivityLogRepository {
	return &activitylogRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *activitylogRepository) CreateMany(c context.Context, items []domain.ActivityLog) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *activitylogRepository) Create(c context.Context, activitylog *domain.ActivityLog) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, activitylog)

	return err
}

func (ur *activitylogRepository) Update(c context.Context, activitylog *domain.ActivityLog) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": activitylog.ID}
	update := bson.M{
		"$set": activitylog,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *activitylogRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *activitylogRepository) Fetch(c context.Context) ([]domain.ActivityLog, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var activitylogs []domain.ActivityLog

	err = cursor.All(c, &activitylogs)
	if activitylogs == nil {
		return []domain.ActivityLog{}, err
	}

	return activitylogs, err
}

func (tr *activitylogRepository) FetchByID(c context.Context, activitylogID string) (domain.ActivityLog, error) {
	collection := tr.database.Collection(tr.collection)

	var activitylog domain.ActivityLog

	idHex, err := primitive.ObjectIDFromHex(activitylogID)
	if err != nil {
		return activitylog, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&activitylog)
	return activitylog, err
}
