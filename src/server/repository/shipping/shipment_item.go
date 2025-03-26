package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type shipmentitemRepository struct {
	database   mongo.Database
	collection string
}

func NewShipmentItemRepository(db mongo.Database, collection string) domain.ShipmentItemRepository {
	return &shipmentitemRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *shipmentitemRepository) CreateMany(c context.Context, items []domain.ShipmentItem) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *shipmentitemRepository) Create(c context.Context, shipmentitem *domain.ShipmentItem) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, shipmentitem)

	return err
}

func (ur *shipmentitemRepository) Update(c context.Context, shipmentitem *domain.ShipmentItem) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": shipmentitem.ID}
	update := bson.M{
		"$set": shipmentitem,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *shipmentitemRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *shipmentitemRepository) Fetch(c context.Context) ([]domain.ShipmentItem, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var shipmentitems []domain.ShipmentItem

	err = cursor.All(c, &shipmentitems)
	if shipmentitems == nil {
		return []domain.ShipmentItem{}, err
	}

	return shipmentitems, err
}

func (tr *shipmentitemRepository) FetchByID(c context.Context, shipmentitemID string) (domain.ShipmentItem, error) {
	collection := tr.database.Collection(tr.collection)

	var shipmentitem domain.ShipmentItem

	idHex, err := primitive.ObjectIDFromHex(shipmentitemID)
	if err != nil {
		return shipmentitem, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&shipmentitem)
	return shipmentitem, err
}
