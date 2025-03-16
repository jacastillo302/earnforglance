package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionWarehouseInventory = "product_inventory_warehouses"
)

// ProductWarehouseInventory represents a record to manage product inventory per warehouse
type ProductWarehouseInventory struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductID        int                `bson:"product_id"`
	WarehouseID      int                `bson:"warehouse_id"`
	StockQuantity    int                `bson:"stock_quantity"`
	ReservedQuantity int                `bson:"reserved_quantity"`
}
