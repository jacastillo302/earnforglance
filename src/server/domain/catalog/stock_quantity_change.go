package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStockQuantityChange = "stock_quantity_changes"
)

// StockQuantityChange represents a stock quantity change entry
type StockQuantityChange struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	QuantityAdjustment int                `bson:"quantity_adjustment"`
	StockQuantity      int                `bson:"stock_quantity"`
	Message            string             `bson:"message"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	ProductID          primitive.ObjectID `bson:"product_id"`
	CombinationID      *int               `bson:"combination_id,omitempty"`
	WarehouseID        *int               `bson:"warehouse_id,omitempty"`
}

type StockQuantityChangeRepository interface {
	Create(c context.Context, stock_quantity_change *StockQuantityChange) error
	Update(c context.Context, stock_quantity_change *StockQuantityChange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StockQuantityChange, error)
	FetchByID(c context.Context, ID string) (StockQuantityChange, error)
}

type StockQuantityChangeUsecase interface {
	FetchByID(c context.Context, ID string) (StockQuantityChange, error)
	Create(c context.Context, stock_quantity_change *StockQuantityChange) error
	Update(c context.Context, stock_quantity_change *StockQuantityChange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StockQuantityChange, error)
}
