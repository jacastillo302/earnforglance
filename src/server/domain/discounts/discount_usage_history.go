package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountUsageHistory = "discount_usage_histories"
)

// DiscountUsageHistory represents a discount usage history entry
type DiscountUsageHistory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DiscountID   primitive.ObjectID `bson:"discount_id"`
	OrderID      primitive.ObjectID `bson:"order_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}

type DiscountUsageHistoryRepository interface {
	CreateMany(c context.Context, items []DiscountUsageHistory) error
	Create(c context.Context, discount_usage_history *DiscountUsageHistory) error
	Update(c context.Context, discount_usage_history *DiscountUsageHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountUsageHistory, error)
	FetchByID(c context.Context, ID string) (DiscountUsageHistory, error)
}

type DiscountUsageHistoryUsecase interface {
	CreateMany(c context.Context, items []DiscountUsageHistory) error
	FetchByID(c context.Context, ID string) (DiscountUsageHistory, error)
	Create(c context.Context, discount_usage_history *DiscountUsageHistory) error
	Update(c context.Context, discount_usage_history *DiscountUsageHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountUsageHistory, error)
}
