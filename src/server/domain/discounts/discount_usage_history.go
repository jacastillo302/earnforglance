package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountUsageHistory = "discount_usage_histories"
)

// DiscountUsageHistory represents a discount usage history entry
type DiscountUsageHistory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	DiscountID   int                `bson:"discount_id"`
	OrderID      int                `bson:"order_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
