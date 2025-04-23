package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (ur *pickupPointRepository) CreateMany(c context.Context, items []domain.PickupPoint) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
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

func (ur *pickupPointRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *pickupPointRepository) Fetch(c context.Context) ([]domain.PickupPoint, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
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

	idHex, err := bson.ObjectIDFromHex(pickupPointID)
	if err != nil {
		return pickupPoint, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&pickupPoint)
	return pickupPoint, err
}
