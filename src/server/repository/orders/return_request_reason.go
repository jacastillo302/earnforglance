package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type returnrequestreasonRepository struct {
	database   mongo.Database
	collection string
}

func NewReturnRequestReasonRepository(db mongo.Database, collection string) domain.ReturnRequestReasonRepository {
	return &returnrequestreasonRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *returnrequestreasonRepository) Create(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, returnrequestreason)

	return err
}

func (ur *returnrequestreasonRepository) Update(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": returnrequestreason.ID}
	update := bson.M{
		"$set": returnrequestreason,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *returnrequestreasonRepository) Delete(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": returnrequestreason.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *returnrequestreasonRepository) Fetch(c context.Context) ([]domain.ReturnRequestReason, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var returnrequestreasons []domain.ReturnRequestReason

	err = cursor.All(c, &returnrequestreasons)
	if returnrequestreasons == nil {
		return []domain.ReturnRequestReason{}, err
	}

	return returnrequestreasons, err
}

func (tr *returnrequestreasonRepository) FetchByID(c context.Context, returnrequestreasonID string) (domain.ReturnRequestReason, error) {
	collection := tr.database.Collection(tr.collection)

	var returnrequestreason domain.ReturnRequestReason

	idHex, err := primitive.ObjectIDFromHex(returnrequestreasonID)
	if err != nil {
		return returnrequestreason, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&returnrequestreason)
	return returnrequestreason, err
}
