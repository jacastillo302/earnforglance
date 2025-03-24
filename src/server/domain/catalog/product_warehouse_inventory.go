package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductWarehouseInventory = "product_warehouse_inventories"
)

// ProductWarehouseInventory represents a record to manage product inventory per warehouse
type ProductWarehouseInventory struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductID        primitive.ObjectID `bson:"product_id"`
	WarehouseID      primitive.ObjectID `bson:"warehouse_id"`
	StockQuantity    int                `bson:"stock_quantity"`
	ReservedQuantity int                `bson:"reserved_quantity"`
}

type ProductWarehouseInventoryRepository interface {
	CreateMany(c context.Context, items []ProductWarehouseInventory) error
	Create(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Update(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductWarehouseInventory, error)
	FetchByID(c context.Context, ID string) (ProductWarehouseInventory, error)
}

type ProductWarehouseInventoryUsecase interface {
	FetchByID(c context.Context, ID string) (ProductWarehouseInventory, error)
	Create(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Update(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductWarehouseInventory, error)
}
