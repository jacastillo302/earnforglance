package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type shipmentRepository struct {
	database   mongo.Database
	collection string
}

func NewShipmentRepository(db mongo.Database, collection string) domain.ShipmentRepository {
	return &shipmentRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shipmentRepository) CreateMany(c context.Context, items []domain.Shipment) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shipmentRepository) Create(c context.Context, shipment *domain.Shipment) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shipment)

	return err
}

func (ur *shipmentRepository) Update(c context.Context, shipment *domain.Shipment) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shipment.ID}
	update := bson.M{
		"$set": shipment,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shipmentRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := bson.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}
func (ur *shipmentRepository) Fetch(c context.Context) ([]domain.Shipment, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().
		SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shipments []domain.Shipment

	err = cursor.All(c, &shipments)
	if shipments == nil {
		return []domain.Shipment{}, err
	}

	return shipments, err
}

func (tr *shipmentRepository) FetchByID(c context.Context, shipmentID string) (domain.Shipment, error) {
	collection := tr.database.Collection(tr.collection)

	var shipment domain.Shipment

	idHex, err := bson.ObjectIDFromHex(shipmentID)
	if err != nil {
		return shipment, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shipment)
	return shipment, err
}
