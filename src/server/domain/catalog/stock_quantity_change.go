package domain

import (
	"context"
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
	ProductID          primitive.ObjectID `bson:"product_id"`
	CombinationID      *int               `bson:"combination_id,omitempty"`
	WarehouseID        *int               `bson:"warehouse_id,omitempty"`
}

type StockQuantityHistoryRepository interface {
	Create(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Update(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Delete(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Fetch(c context.Context) ([]StockQuantityHistory, error)
	FetchByID(c context.Context, stock_quantity_changeID string) (StockQuantityHistory, error)
}

type StockQuantityHistoryUsecase interface {
	FetchByID(c context.Context, stock_quantity_changeID string) (StockQuantityHistory, error)
	Create(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Update(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Delete(c context.Context, stock_quantity_change *StockQuantityHistory) error
	Fetch(c context.Context) ([]StockQuantityHistory, error)
}
