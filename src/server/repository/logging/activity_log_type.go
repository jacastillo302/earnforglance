package repository

import (
	"context"

	domain "earnforglance/server/domain/logging"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type activitylogtypeRepository struct {
	database   mongo.Database
	collection string
}

func NewActivityLogTypeRepository(db mongo.Database, collection string) domain.ActivityLogTypeRepository {
	return &activitylogtypeRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *activitylogtypeRepository) CreateMany(c context.Context, items []domain.ActivityLogType) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *activitylogtypeRepository) Create(c context.Context, activitylogtype *domain.ActivityLogType) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, activitylogtype)

	return err
}

func (ur *activitylogtypeRepository) Update(c context.Context, activitylogtype *domain.ActivityLogType) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": activitylogtype.ID}
	update := bson.M{
		"$set": activitylogtype,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *activitylogtypeRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *activitylogtypeRepository) Fetch(c context.Context) ([]domain.ActivityLogType, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var activitylogtypes []domain.ActivityLogType

	err = cursor.All(c, &activitylogtypes)
	if activitylogtypes == nil {
		return []domain.ActivityLogType{}, err
	}

	return activitylogtypes, err
}

func (tr *activitylogtypeRepository) FetchByID(c context.Context, activitylogtypeID string) (domain.ActivityLogType, error) {
	collection := tr.database.Collection(tr.collection)

	var activitylogtype domain.ActivityLogType

	idHex, err := bson.ObjectIDFromHex(activitylogtypeID)
	if err != nil {
		return activitylogtype, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&activitylogtype)
	return activitylogtype, err
}
