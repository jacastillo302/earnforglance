package repository

import (
	"context"

	domain "earnforglance/server/domain/scheduleTasks"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type scheduletaskRepository struct {
	database   mongo.Database
	collection string
}

func NewScheduleTaskRepository(db mongo.Database, collection string) domain.ScheduleTaskRepository {
	return &scheduletaskRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *scheduletaskRepository) CreateMany(c context.Context, items []domain.ScheduleTask) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *scheduletaskRepository) Create(c context.Context, scheduletask *domain.ScheduleTask) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, scheduletask)

	return err
}

func (ur *scheduletaskRepository) Update(c context.Context, scheduletask *domain.ScheduleTask) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": scheduletask.ID}
	update := bson.M{
		"$set": scheduletask,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *scheduletaskRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *scheduletaskRepository) Fetch(c context.Context) ([]domain.ScheduleTask, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var scheduletasks []domain.ScheduleTask

	err = cursor.All(c, &scheduletasks)
	if scheduletasks == nil {
		return []domain.ScheduleTask{}, err
	}

	return scheduletasks, err
}

func (tr *scheduletaskRepository) FetchByID(c context.Context, scheduletaskID string) (domain.ScheduleTask, error) {
	collection := tr.database.Collection(tr.collection)

	var scheduletask domain.ScheduleTask

	idHex, err := bson.ObjectIDFromHex(scheduletaskID)
	if err != nil {
		return scheduletask, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&scheduletask)
	return scheduletask, err
}
