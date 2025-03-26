package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type privatemessageRepository struct {
	database   mongo.Database
	collection string
}

func NewPrivateMessageRepository(db mongo.Database, collection string) domain.PrivateMessageRepository {
	return &privatemessageRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *privatemessageRepository) CreateMany(c context.Context, items []domain.PrivateMessage) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *privatemessageRepository) Create(c context.Context, privatemessage *domain.PrivateMessage) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, privatemessage)

	return err
}

func (ur *privatemessageRepository) Update(c context.Context, privatemessage *domain.PrivateMessage) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": privatemessage.ID}
	update := bson.M{
		"$set": privatemessage,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *privatemessageRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *privatemessageRepository) Fetch(c context.Context) ([]domain.PrivateMessage, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var privatemessages []domain.PrivateMessage

	err = cursor.All(c, &privatemessages)
	if privatemessages == nil {
		return []domain.PrivateMessage{}, err
	}

	return privatemessages, err
}

func (tr *privatemessageRepository) FetchByID(c context.Context, privatemessageID string) (domain.PrivateMessage, error) {
	collection := tr.database.Collection(tr.collection)

	var privatemessage domain.PrivateMessage

	idHex, err := primitive.ObjectIDFromHex(privatemessageID)
	if err != nil {
		return privatemessage, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&privatemessage)
	return privatemessage, err
}
