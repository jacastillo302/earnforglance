package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBackInStockSubscription = "back_in_stock_subscriptions"
)

// BackInStockSubscription represents a back in stock subscription
type BackInStockSubscription struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	StoreID      primitive.ObjectID `bson:"store_id"`
	ProductID    primitive.ObjectID `bson:"product_id"`
	CustomerID   primitive.ObjectID `bson:"customer_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}

type BackInStockSubscriptionRepository interface {
	CreateMany(c context.Context, items []BackInStockSubscription) error
	Create(c context.Context, back_in_stock_subscription *BackInStockSubscription) error
	Update(c context.Context, back_in_stock_subscription *BackInStockSubscription) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BackInStockSubscription, error)
	FetchByID(c context.Context, ID string) (BackInStockSubscription, error)
}

type BackInStockSubscriptionUsecase interface {
	CreateMany(c context.Context, items []BackInStockSubscription) error
	FetchByID(c context.Context, ID string) (BackInStockSubscription, error)
	Create(c context.Context, back_in_stock_subscription *BackInStockSubscription) error
	Update(c context.Context, back_in_stock_subscription *BackInStockSubscription) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BackInStockSubscription, error)
}
