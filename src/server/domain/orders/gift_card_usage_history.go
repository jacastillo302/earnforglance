package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGiftCardUsageHistory = "gift_card_usage_histories"
)

// GiftCardUsageHistory represents a gift card usage history entry
type GiftCardUsageHistory struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	GiftCardID      int                `bson:"gift_card_id"`
	UsedWithOrderID int                `bson:"used_with_order_id"`
	UsedValue       float64            `bson:"used_value"`
	CreatedOnUtc    time.Time          `bson:"created_on_utc"`
}
