package repository

import (
	"context"

	domain "earnforglance/server/domain/polls"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type pollRepository struct {
	database   mongo.Database
	collection string
}

func NewPollRepository(db mongo.Database, collection string) domain.PollRepository {
	return &pollRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pollRepository) CreateMany(c context.Context, items []domain.Poll) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *pollRepository) Create(c context.Context, poll *domain.Poll) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, poll)

	return err
}

func (ur *pollRepository) Update(c context.Context, poll *domain.Poll) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": poll.ID}
	update := bson.M{
		"$set": poll,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *pollRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pollRepository) Fetch(c context.Context) ([]domain.Poll, error) {
	collection := ur.database.Collection(ur.collection)

	findOptions := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})

	cursor, err := collection.Find(c, bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	var polls []domain.Poll

	err = cursor.All(c, &polls)
	if polls == nil {
		return []domain.Poll{}, err
	}

	return polls, err
}

func (tr *pollRepository) FetchByID(c context.Context, pollID string) (domain.Poll, error) {
	collection := tr.database.Collection(tr.collection)

	var poll domain.Poll

	idHex, err := bson.ObjectIDFromHex(pollID)
	if err != nil {
		return poll, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&poll)
	return poll, err
}
