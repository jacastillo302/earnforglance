package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type messageTemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewMessageTemplateRepository(db mongo.Database, collection string) domain.MessageTemplateRepository {
	return &messageTemplateRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *messageTemplateRepository) CreateMany(c context.Context, items []domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *messageTemplateRepository) Create(c context.Context, MessageTemplate *domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, MessageTemplate)

	return err
}

func (ur *messageTemplateRepository) Update(c context.Context, MessageTemplate *domain.MessageTemplate) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": MessageTemplate.ID}
	update := bson.M{
		"$set": MessageTemplate,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *messageTemplateRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *messageTemplateRepository) Fetch(c context.Context) ([]domain.MessageTemplate, error) {
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

func (tr *messageTemplateRepository) FetchByID(c context.Context, MessageTemplateID string) (domain.MessageTemplate, error) {
	collection := tr.database.Collection(tr.collection)

	var MessageTemplate domain.MessageTemplate

	idHex, err := primitive.ObjectIDFromHex(MessageTemplateID)
	if err != nil {
		return MessageTemplate, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&MessageTemplate)
	return MessageTemplate, err
}
