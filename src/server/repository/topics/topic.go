package repository

import (
	"context"

	domain "earnforglance/server/domain/topics"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type topicRepository struct {
	database   mongo.Database
	collection string
}

func NewTopicRepository(db mongo.Database, collection string) domain.TopicRepository {
	return &topicRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *topicRepository) CreateMany(c context.Context, items []domain.Topic) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *topicRepository) Create(c context.Context, topic *domain.Topic) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, topic)

	return err
}

func (ur *topicRepository) Update(c context.Context, topic *domain.Topic) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": topic.ID}
	update := bson.M{
		"$set": topic,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *topicRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *topicRepository) Fetch(c context.Context) ([]domain.Topic, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var topics []domain.Topic

	err = cursor.All(c, &topics)
	if topics == nil {
		return []domain.Topic{}, err
	}

	return topics, err
}

func (tr *topicRepository) FetchByID(c context.Context, topicID string) (domain.Topic, error) {
	collection := tr.database.Collection(tr.collection)

	var topic domain.Topic

	idHex, err := bson.ObjectIDFromHex(topicID)
	if err != nil {
		return topic, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&topic)
	return topic, err
}
