package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTierPrice = "tier_prices"
)

// TierPrice represents a tier price
type TierPrice struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductID        int                `bson:"product_id"`
	StoreID          int                `bson:"store_id"`
	CustomerRoleID   *int               `bson:"customer_role_id,omitempty"`
	Quantity         int                `bson:"quantity"`
	Price            float64            `bson:"price"`
	StartDateTimeUtc *time.Time         `bson:"start_date_time_utc,omitempty"`
	EndDateTimeUtc   *time.Time         `bson:"end_date_time_utc,omitempty"`
}
