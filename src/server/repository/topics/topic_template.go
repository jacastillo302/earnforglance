package repository

import (
	"context"

	domain "earnforglance/server/domain/topics"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type topictemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewTopicTemplateRepository(db mongo.Database, collection string) domain.TopicTemplateRepository {
	return &topictemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *topictemplateRepository) CreateMany(c context.Context, items []domain.TopicTemplate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *topictemplateRepository) Create(c context.Context, topictemplate *domain.TopicTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, topictemplate)

	return err
}

func (ur *topictemplateRepository) Update(c context.Context, topictemplate *domain.TopicTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": topictemplate.ID}
	update := bson.M{
		"$set": topictemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *topictemplateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *topictemplateRepository) Fetch(c context.Context) ([]domain.TopicTemplate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var topictemplates []domain.TopicTemplate

	err = cursor.All(c, &topictemplates)
	if topictemplates == nil {
		return []domain.TopicTemplate{}, err
	}

	return topictemplates, err
}

func (tr *topictemplateRepository) FetchByID(c context.Context, topictemplateID string) (domain.TopicTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var topictemplate domain.TopicTemplate

	idHex, err := bson.ObjectIDFromHex(topictemplateID)
	if err != nil {
		return topictemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&topictemplate)
	return topictemplate, err
}
