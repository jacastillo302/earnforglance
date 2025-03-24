package repository

import (
	"context"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type warehouseRepository struct {
	database   mongo.Database
	collection string
}

func NewWarehouseRepository(db mongo.Database, collection string) domain.WarehouseRepository {
	return &warehouseRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *warehouseRepository) CreateMany(c context.Context, items []domain.Warehouse) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *warehouseRepository) Create(c context.Context, warehouse *domain.Warehouse) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, warehouse)

	return err
}

func (ur *warehouseRepository) Update(c context.Context, warehouse *domain.Warehouse) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"_id": warehouse.ID}
	update := bson.M{
		"$set": warehouse,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (ur *warehouseRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *warehouseRepository) Fetch(c context.Context) ([]domain.Warehouse, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var warehouses []domain.Warehouse

	err = cursor.All(c, &warehouses)
	if warehouses == nil {
		return []domain.Warehouse{}, err
	}

	return warehouses, err
}

func (tr *warehouseRepository) FetchByID(c context.Context, warehouseID string) (domain.Warehouse, error) {
	collection := tr.database.Collection(tr.collection)

	var warehouse domain.Warehouse

	idHex, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return warehouse, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&warehouse)
	return warehouse, err
}
