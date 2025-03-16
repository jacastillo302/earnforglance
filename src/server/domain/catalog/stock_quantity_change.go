package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStockQuantityHistory = "stock_quantity_changes"
)

// StockQuantityHistory represents a stock quantity change entry
type StockQuantityHistory struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	QuantityAdjustment int                `bson:"quantity_adjustment"`
	StockQuantity      int                `bson:"stock_quantity"`
	Message            string             `bson:"message"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	ProductID          int                `bson:"product_id"`
	CombinationID      *int               `bson:"combination_id,omitempty"`
	WarehouseID        *int               `bson:"warehouse_id,omitempty"`
}
