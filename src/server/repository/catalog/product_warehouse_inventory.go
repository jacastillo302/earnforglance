package repository

import (
	"context"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/service/data/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productwarehouseinventoryRepository struct {
	database   mongo.Database
	collection string
}

func NewProductWarehouseInventoryRepository(db mongo.Database, collection string) domain.ProductWarehouseInventoryRepository {
	return &productwarehouseinventoryRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *productwarehouseinventoryRepository) CreateMany(c context.Context, items []domain.ProductWarehouseInventory) error {
	collection := ur.database.Collection(ur.collection)

	interfaces := make([]interface{}, len(items))
	for i, item := range items {
		interfaces[i] = item
	}

	_, err := collection.InsertMany(c, interfaces)

	return err
}

func (ur *productwarehouseinventoryRepository) Create(c context.Context, productwarehouseinventory *domain.ProductWarehouseInventory) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, productwarehouseinventory)

	return err
}

func (ur *productwarehouseinventoryRepository) Update(c context.Context, productwarehouseinventory *domain.ProductWarehouseInventory) error {
	collection := ur.database.Collection(ur.collection)

	filter := bson.M{"warehouse_id": productwarehouseinventory.WarehouseID, "product_id": productwarehouseinventory.ProductID}
	update := bson.M{
		"$set": productwarehouseinventory,
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err

}

func (ur *productwarehouseinventoryRepository) Delete(c context.Context, ID string) error {
	collection := ur.database.Collection(ur.collection)

	idHex, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})

	return err

}

func (ur *productwarehouseinventoryRepository) Fetch(c context.Context) ([]domain.ProductWarehouseInventory, error) {
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var productwarehouseinventorys []domain.ProductWarehouseInventory

	err = cursor.All(c, &productwarehouseinventorys)
	if productwarehouseinventorys == nil {
		return []domain.ProductWarehouseInventory{}, err
	}

	return productwarehouseinventorys, err
}

func (tr *productwarehouseinventoryRepository) FetchByID(c context.Context, productwarehouseinventoryID string) (domain.ProductWarehouseInventory, error) {
	collection := tr.database.Collection(tr.collection)

	var productwarehouseinventory domain.ProductWarehouseInventory

	idHex, err := primitive.ObjectIDFromHex(productwarehouseinventoryID)
	if err != nil {
		return productwarehouseinventory, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&productwarehouseinventory)
	return productwarehouseinventory, err
}
