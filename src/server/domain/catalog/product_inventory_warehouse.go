package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionWarehouseInventory = "product_inventory_warehouses"
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
	Create(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Update(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Delete(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Fetch(c context.Context) ([]ProductWarehouseInventory, error)
	FetchByID(c context.Context, product_inventory_warehouseID string) (ProductWarehouseInventory, error)
}

type ProductWarehouseInventoryUsecase interface {
	FetchByID(c context.Context, product_inventory_warehouseID string) (ProductWarehouseInventory, error)
	Create(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Update(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Delete(c context.Context, product_inventory_warehouse *ProductWarehouseInventory) error
	Fetch(c context.Context) ([]ProductWarehouseInventory, error)
}
