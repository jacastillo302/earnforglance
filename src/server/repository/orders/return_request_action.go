package repository

import (
	"context"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type returnrequestactionRepository struct {
	database   mongo.Database
	collection string
}

func NewReturnRequestActionRepository(db mongo.Database, collection string) domain.ReturnRequestActionRepository {
	return &returnrequestactionRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *returnrequestactionRepository) CreateMany(c context.Context, items []domain.ReturnRequestAction) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *returnrequestactionRepository) Create(c context.Context, returnrequestaction *domain.ReturnRequestAction) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, returnrequestaction)

	return err
}

func (ur *returnrequestactionRepository) Update(c context.Context, returnrequestaction *domain.ReturnRequestAction) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": returnrequestaction.ID}
	update := bson.M{
		"$set": returnrequestaction,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *returnrequestactionRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *returnrequestactionRepository) Fetch(c context.Context) ([]domain.ReturnRequestAction, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var returnrequestactions []domain.ReturnRequestAction

	err = cursor.All(c, &returnrequestactions)
	if returnrequestactions == nil {
		return []domain.ReturnRequestAction{}, err
	}

	return returnrequestactions, err
}

func (tr *returnrequestactionRepository) FetchByID(c context.Context, returnrequestactionID string) (domain.ReturnRequestAction, error) {
	collection := tr.database.Collection(tr.collection)

	var returnrequestaction domain.ReturnRequestAction

	idHex, err := bson.ObjectIDFromHex(returnrequestactionID)
	if err != nil {
		return returnrequestaction, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&returnrequestaction)
	return returnrequestaction, err
}
