package repository

import (
	"context"

	domain "earnforglance/server/domain/forums"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type forumsubscriptionRepository struct {
	database   mongo.Database
	collection string
}

func NewForumSubscriptionRepository(db mongo.Database, collection string) domain.ForumSubscriptionRepository {
	return &forumsubscriptionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *forumsubscriptionRepository) Create(c context.Context, forumsubscription *domain.ForumSubscription) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, forumsubscription)

	return err
}

func (ur *forumsubscriptionRepository) Update(c context.Context, forumsubscription *domain.ForumSubscription) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": forumsubscription.ID}
	update := bson.M{
		"$set": forumsubscription,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *forumsubscriptionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *forumsubscriptionRepository) Fetch(c context.Context) ([]domain.ForumSubscription, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var forumsubscriptions []domain.ForumSubscription

	err = cursor.All(c, &forumsubscriptions)
	if forumsubscriptions == nil {
		return []domain.ForumSubscription{}, err
	}

	return forumsubscriptions, err
}

func (tr *forumsubscriptionRepository) FetchByID(c context.Context, forumsubscriptionID string) (domain.ForumSubscription, error) {
	collection := tr.database.Collection(tr.collection)

	var forumsubscription domain.ForumSubscription

	idHex, err := primitive.ObjectIDFromHex(forumsubscriptionID)
	if err != nil {
		return forumsubscription, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&forumsubscription)
	return forumsubscription, err
}
