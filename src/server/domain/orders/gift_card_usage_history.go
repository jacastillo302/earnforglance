package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionGiftCardUsageHistory = "gift_card_usage_histories"
)

// GiftCardUsageHistory represents a gift card usage history entry
type GiftCardUsageHistory struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	GiftCardID   bson.ObjectID `bson:"gift_card_id"`
	OrderID      bson.ObjectID `bson:"used_with_order_id"`
	UsedValue    float64       `bson:"used_value"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

// GiftCardUsageHistoryRepository defines the repository interface for GiftCardUsageHistory
type GiftCardUsageHistoryRepository interface {
	CreateMany(c context.Context, items []GiftCardUsageHistory) error
	Create(c context.Context, gift_card_usage_history *GiftCardUsageHistory) error
	Update(c context.Context, gift_card_usage_history *GiftCardUsageHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCardUsageHistory, error)
	FetchByID(c context.Context, ID string) (GiftCardUsageHistory, error)
}

// GiftCardUsageHistoryUsecase defines the usecase interface for GiftCardUsageHistory
type GiftCardUsageHistoryUsecase interface {
	CreateMany(c context.Context, items []GiftCardUsageHistory) error
	FetchByID(c context.Context, ID string) (GiftCardUsageHistory, error)
	Create(c context.Context, gift_card_usage_history *GiftCardUsageHistory) error
	Update(c context.Context, gift_card_usage_history *GiftCardUsageHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GiftCardUsageHistory, error)
}
