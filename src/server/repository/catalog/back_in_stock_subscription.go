package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type backinstocksubscriptionRepository struct {
	database   mongo.Database
	collection string
}

func NewBackInStockSubscriptionRepository(db mongo.Database, collection string) domain.BackInStockSubscriptionRepository {
	return &backinstocksubscriptionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *backinstocksubscriptionRepository) Create(c context.Context, backinstocksubscription *domain.BackInStockSubscription) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, backinstocksubscription)

	return err
}

func (ur *backinstocksubscriptionRepository) Update(c context.Context, backinstocksubscription *domain.BackInStockSubscription) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": backinstocksubscription.ID}
	update := bson.M{
		"$set": backinstocksubscription,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *backinstocksubscriptionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *backinstocksubscriptionRepository) Fetch(c context.Context) ([]domain.BackInStockSubscription, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var backinstocksubscriptions []domain.BackInStockSubscription

	err = cursor.All(c, &backinstocksubscriptions)
	if backinstocksubscriptions == nil {
		return []domain.BackInStockSubscription{}, err
	}

	return backinstocksubscriptions, err
}

func (tr *backinstocksubscriptionRepository) FetchByID(c context.Context, backinstocksubscriptionID string) (domain.BackInStockSubscription, error) {
	collection := tr.database.Collection(tr.collection)

	var backinstocksubscription domain.BackInStockSubscription

	idHex, err := primitive.ObjectIDFromHex(backinstocksubscriptionID)
	if err != nil {
		return backinstocksubscription, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&backinstocksubscription)
	return backinstocksubscription, err
}
