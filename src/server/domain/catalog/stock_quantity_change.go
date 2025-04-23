package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionStockQuantityChange = "stock_quantity_changes"
)

// StockQuantityChange represents a stock quantity change entry
type StockQuantityChange struct {
	ID                 bson.ObjectID  `bson:"_id,omitempty"`
	QuantityAdjustment int            `bson:"quantity_adjustment"`
	StockQuantity      int            `bson:"stock_quantity"`
	Message            string         `bson:"message"`
	CreatedOnUtc       time.Time      `bson:"created_on_utc"`
	ProductID          bson.ObjectID  `bson:"product_id"`
	CombinationID      *bson.ObjectID `bson:"combination_id"`
	WarehouseID        bson.ObjectID  `bson:"warehouse_id"`
}

type StockQuantityChangeRepository interface {
	CreateMany(c context.Context, items []StockQuantityChange) error
	Create(c context.Context, stock_quantity_change *StockQuantityChange) error
	Update(c context.Context, stock_quantity_change *StockQuantityChange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StockQuantityChange, error)
	FetchByID(c context.Context, ID string) (StockQuantityChange, error)
}

type StockQuantityChangeUsecase interface {
	CreateMany(c context.Context, items []StockQuantityChange) error
	FetchByID(c context.Context, ID string) (StockQuantityChange, error)
	Create(c context.Context, stock_quantity_change *StockQuantityChange) error
	Update(c context.Context, stock_quantity_change *StockQuantityChange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StockQuantityChange, error)
}
