package repository

import (
	"context"

	domain "earnforglance/server/domain/gdpr"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customerpermanentlydeletedRepository struct {
	database   mongo.Database
	collection string
}

func NewCustomerPermanentlyDeletedRepository(db mongo.Database, collection string) domain.CustomerPermanentlyDeletedRepository {
	return &customerpermanentlydeletedRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *customerpermanentlydeletedRepository) Create(c context.Context, customerpermanentlydeleted *domain.CustomerPermanentlyDeleted) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, customerpermanentlydeleted)

	return err
}

func (ur *customerpermanentlydeletedRepository) Update(c context.Context, customerpermanentlydeleted *domain.CustomerPermanentlyDeleted) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": customerpermanentlydeleted.CustomerID}
	update := bson.M{
		"$set": customerpermanentlydeleted,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *customerpermanentlydeletedRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *customerpermanentlydeletedRepository) Fetch(c context.Context) ([]domain.CustomerPermanentlyDeleted, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var customerpermanentlydeleteds []domain.CustomerPermanentlyDeleted

	err = cursor.All(c, &customerpermanentlydeleteds)
	if customerpermanentlydeleteds == nil {
		return []domain.CustomerPermanentlyDeleted{}, err
	}

	return customerpermanentlydeleteds, err
}

func (tr *customerpermanentlydeletedRepository) FetchByID(c context.Context, customerpermanentlydeletedID string) (domain.CustomerPermanentlyDeleted, error) {
	collection := tr.database.Collection(tr.collection)

	var customerpermanentlydeleted domain.CustomerPermanentlyDeleted

	idHex, err := primitive.ObjectIDFromHex(customerpermanentlydeletedID)
	if err != nil {
		return customerpermanentlydeleted, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&customerpermanentlydeleted)
	return customerpermanentlydeleted, err
}
