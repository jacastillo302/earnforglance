package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type forumpostvoteRepository struct {
	database   mongo.Database
	collection string
}

func NewForumPostVoteRepository(db mongo.Database, collection string) domain.ForumPostVoteRepository {
	return &forumpostvoteRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumpostvoteRepository) CreateMany(c context.Context, items []domain.ForumPostVote) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *forumpostvoteRepository) Create(c context.Context, forumpostvote *domain.ForumPostVote) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumpostvote)

	return err
}

func (ur *forumpostvoteRepository) Update(c context.Context, forumpostvote *domain.ForumPostVote) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumpostvote.ID}
	update := bson.M{
		"$set": forumpostvote,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumpostvoteRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumpostvoteRepository) Fetch(c context.Context) ([]domain.ForumPostVote, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumpostvotes []domain.ForumPostVote

	err = cursor.All(c, &forumpostvotes)
	if forumpostvotes == nil {
		return []domain.ForumPostVote{}, err
	}

	return forumpostvotes, err
}

func (tr *forumpostvoteRepository) FetchByID(c context.Context, forumpostvoteID string) (domain.ForumPostVote, error) {
	collection := tr.database.Collection(tr.collection)

	var forumpostvote domain.ForumPostVote

	idHex, err := bson.ObjectIDFromHex(forumpostvoteID)
	if err != nil {
		return forumpostvote, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumpostvote)
	return forumpostvote, err
}
