package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type returnrequestRepository struct {
	database   mongo.Database
	collection string
}

func NewReturnRequestRepository(db mongo.Database, collection string) domain.ReturnRequestRepository {
	return &returnrequestRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *returnrequestRepository) Create(c context.Context, returnrequest *domain.ReturnRequest) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, returnrequest)

	return err
}

func (ur *returnrequestRepository) Update(c context.Context, returnrequest *domain.ReturnRequest) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": returnrequest.ID}
	update := bson.M{
		"$set": returnrequest,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *returnrequestRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *returnrequestRepository) Fetch(c context.Context) ([]domain.ReturnRequest, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var returnrequests []domain.ReturnRequest

	err = cursor.All(c, &returnrequests)
	if returnrequests == nil {
		return []domain.ReturnRequest{}, err
	}

	return returnrequests, err
}

func (tr *returnrequestRepository) FetchByID(c context.Context, returnrequestID string) (domain.ReturnRequest, error) {
	collection := tr.database.Collection(tr.collection)

	var returnrequest domain.ReturnRequest

	idHex, err := primitive.ObjectIDFromHex(returnrequestID)
	if err != nil {
		return returnrequest, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&returnrequest)
	return returnrequest, err
}
