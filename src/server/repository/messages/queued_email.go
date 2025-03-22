package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type queuedemailRepository struct {
	database   mongo.Database
	collection string
}

func NewQueuedEmailRepository(db mongo.Database, collection string) domain.QueuedEmailRepository {
	return &queuedemailRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *queuedemailRepository) Create(c context.Context, queuedemail *domain.QueuedEmail) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, queuedemail)

	return err
}

func (ur *queuedemailRepository) Update(c context.Context, queuedemail *domain.QueuedEmail) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": queuedemail.ID}
	update := bson.M{
		"$set": queuedemail,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *queuedemailRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *queuedemailRepository) Fetch(c context.Context) ([]domain.QueuedEmail, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var queuedemails []domain.QueuedEmail

	err = cursor.All(c, &queuedemails)
	if queuedemails == nil {
		return []domain.QueuedEmail{}, err
	}

	return queuedemails, err
}

func (tr *queuedemailRepository) FetchByID(c context.Context, queuedemailID string) (domain.QueuedEmail, error) {
	collection := tr.database.Collection(tr.collection)

	var queuedemail domain.QueuedEmail

	idHex, err := primitive.ObjectIDFromHex(queuedemailID)
	if err != nil {
		return queuedemail, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&queuedemail)
	return queuedemail, err
}
