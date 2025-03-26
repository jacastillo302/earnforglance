package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type forumtopicRepository struct {
	database   mongo.Database
	collection string
}

func NewForumTopicRepository(db mongo.Database, collection string) domain.ForumTopicRepository {
	return &forumtopicRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumtopicRepository) CreateMany(c context.Context, items []domain.ForumTopic) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *forumtopicRepository) Create(c context.Context, forumtopic *domain.ForumTopic) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumtopic)

	return err
}

func (ur *forumtopicRepository) Update(c context.Context, forumtopic *domain.ForumTopic) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumtopic.ID}
	update := bson.M{
		"$set": forumtopic,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumtopicRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumtopicRepository) Fetch(c context.Context) ([]domain.ForumTopic, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumtopics []domain.ForumTopic

	err = cursor.All(c, &forumtopics)
	if forumtopics == nil {
		return []domain.ForumTopic{}, err
	}

	return forumtopics, err
}

func (tr *forumtopicRepository) FetchByID(c context.Context, forumtopicID string) (domain.ForumTopic, error) {
	collection := tr.database.Collection(tr.collection)

	var forumtopic domain.ForumTopic

	idHex, err := primitive.ObjectIDFromHex(forumtopicID)
	if err != nil {
		return forumtopic, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumtopic)
	return forumtopic, err
}
