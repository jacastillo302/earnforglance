package repository

import (
	"context"

	domain "earnforglance/server/domain/polls"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type pollvotingrecordRepository struct {
	database   mongo.Database
	collection string
}

func NewPollVotingRecordRepository(db mongo.Database, collection string) domain.PollVotingRecordRepository {
	return &pollvotingrecordRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pollvotingrecordRepository) CreateMany(c context.Context, items []domain.PollVotingRecord) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *pollvotingrecordRepository) Create(c context.Context, pollvotingrecord *domain.PollVotingRecord) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, pollvotingrecord)

	return err
}

func (ur *pollvotingrecordRepository) Update(c context.Context, pollvotingrecord *domain.PollVotingRecord) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": pollvotingrecord.ID}
	update := bson.M{
		"$set": pollvotingrecord,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *pollvotingrecordRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pollvotingrecordRepository) Fetch(c context.Context) ([]domain.PollVotingRecord, error) {
	collection := ur.database.Collection(ur.collection)

	findOptions := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	var pollvotingrecords []domain.PollVotingRecord

	err = cursor.All(c, &pollvotingrecords)
	if pollvotingrecords == nil {
		return []domain.PollVotingRecord{}, err
	}

	return pollvotingrecords, err
}

func (tr *pollvotingrecordRepository) FetchByID(c context.Context, pollvotingrecordID string) (domain.PollVotingRecord, error) {
	collection := tr.database.Collection(tr.collection)

	var pollvotingrecord domain.PollVotingRecord

	idHex, err := bson.ObjectIDFromHex(pollvotingrecordID)
	if err != nil {
		return pollvotingrecord, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&pollvotingrecord)
	return pollvotingrecord, err
}
