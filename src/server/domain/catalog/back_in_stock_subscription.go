package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBackInStockSubscription = "back_in_stock_subscriptions"
)

// BackInStockSubscription represents a back in stock subscription
type BackInStockSubscription struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	StoreID      int                `bson:"store_id"`
	ProductID    int                `bson:"product_id"`
	CustomerID   int                `bson:"customer_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
