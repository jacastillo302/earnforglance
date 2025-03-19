package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type pickupPointRepository struct {
	database   mongo.Database
	collection string
}

func NewPickupPointRepository(db mongo.Database, collection string) domain.PickupPointRepository {
	return &pickupPointRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *pickupPointRepository) Create(c context.Context, pickupPoint *domain.PickupPoint) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, pickupPoint)

	return err
}

func (ur *pickupPointRepository) Update(c context.Context, pickupPoint *domain.PickupPoint) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": pickupPoint.ID}
	update := bson.M{
		"$set": pickupPoint,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *pickupPointRepository) Delete(c context.Context, pickupPoint *domain.PickupPoint) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": pickupPoint.ID}
	_, err := collection.DeleteOne(c, filter)
	return err
}

func (ur *pickupPointRepository) Fetch(c context.Context) ([]domain.PickupPoint, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var pickupPoints []domain.PickupPoint

	err = cursor.All(c, &pickupPoints)
	if pickupPoints == nil {
		return []domain.PickupPoint{}, err
	}

	return pickupPoints, err
}

func (tr *pickupPointRepository) FetchByID(c context.Context, pickupPointID string) (domain.PickupPoint, error) {
	collection := tr.database.Collection(tr.collection)

	var pickupPoint domain.PickupPoint

	idHex, err := primitive.ObjectIDFromHex(pickupPointID)
	if err != nil {
		return pickupPoint, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&pickupPoint)
	return pickupPoint, err
}
