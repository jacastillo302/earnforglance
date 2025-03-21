package repository

import (
	"context"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type newslettersubscriptionRepository struct {
	database   mongo.Database
	collection string
}

func NewNewsLetterSubscriptionRepository(db mongo.Database, collection string) domain.NewsLetterSubscriptionRepository {
	return &newslettersubscriptionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *newslettersubscriptionRepository) Create(c context.Context, newslettersubscription *domain.NewsLetterSubscription) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, newslettersubscription)

	return err
}

func (ur *newslettersubscriptionRepository) Update(c context.Context, newslettersubscription *domain.NewsLetterSubscription) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": newslettersubscription.ID}
	update := bson.M{
		"$set": newslettersubscription,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *newslettersubscriptionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *newslettersubscriptionRepository) Fetch(c context.Context) ([]domain.NewsLetterSubscription, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var newslettersubscriptions []domain.NewsLetterSubscription

	err = cursor.All(c, &newslettersubscriptions)
	if newslettersubscriptions == nil {
		return []domain.NewsLetterSubscription{}, err
	}

	return newslettersubscriptions, err
}

func (tr *newslettersubscriptionRepository) FetchByID(c context.Context, newslettersubscriptionID string) (domain.NewsLetterSubscription, error) {
	collection := tr.database.Collection(tr.collection)

	var newslettersubscription domain.NewsLetterSubscription

	idHex, err := primitive.ObjectIDFromHex(newslettersubscriptionID)
	if err != nil {
		return newslettersubscription, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&newslettersubscription)
	return newslettersubscription, err
}
