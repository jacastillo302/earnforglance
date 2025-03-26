package repository

import (
	"context"

	domain "earnforglance/server/domain/polls"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pollanswerRepository struct {
	database   mongo.Database
	collection string
}

func NewPollAnswerRepository(db mongo.Database, collection string) domain.PollAnswerRepository {
	return &pollanswerRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pollanswerRepository) CreateMany(c context.Context, items []domain.PollAnswer) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *pollanswerRepository) Create(c context.Context, pollanswer *domain.PollAnswer) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, pollanswer)

	return err
}

func (ur *pollanswerRepository) Update(c context.Context, pollanswer *domain.PollAnswer) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": pollanswer.ID}
	update := bson.M{
		"$set": pollanswer,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *pollanswerRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pollanswerRepository) Fetch(c context.Context) ([]domain.PollAnswer, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var pollanswers []domain.PollAnswer

	err = cursor.All(c, &pollanswers)
	if pollanswers == nil {
		return []domain.PollAnswer{}, err
	}

	return pollanswers, err
}

func (tr *pollanswerRepository) FetchByID(c context.Context, pollanswerID string) (domain.PollAnswer, error) {
	collection := tr.database.Collection(tr.collection)

	var pollanswer domain.PollAnswer

	idHex, err := primitive.ObjectIDFromHex(pollanswerID)
	if err != nil {
		return pollanswer, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&pollanswer)
	return pollanswer, err
}
