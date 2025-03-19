package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageTemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageTemplateRepository(db mongo.Database, collection string) domain.MessageTemplateRepository {
	return &MessageTemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *MessageTemplateRepository) Create(c context.Context, MessageTemplate *domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessageTemplate)

	return err
}

func (ur *MessageTemplateRepository) Update(c context.Context, MessageTemplate *domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplate.ID}
	update := bson.M{
		"$set": MessageTemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *MessageTemplateRepository) Delete(c context.Context, MessageTemplate *domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplate.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *MessageTemplateRepository) Fetch(c context.Context) ([]domain.MessageTemplate, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var MessageTemplates []domain.MessageTemplate

	err = cursor.All(c, &MessageTemplates)
	if MessageTemplates == nil {
		return []domain.MessageTemplate{}, err
	}

	return MessageTemplates, err
}

func (tr *MessageTemplateRepository) FetchByID(c context.Context, MessageTemplateID string) (domain.MessageTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var MessageTemplate domain.MessageTemplate

	idHex, err := primitive.ObjectIDFromHex(MessageTemplateID)
	if err != nil {
		return MessageTemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessageTemplate)
	return MessageTemplate, err
}
