package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionTierPrice = "tier_prices"
)

// TierPrice represents a tier price
type TierPrice struct {
	ID               bson.ObjectID  `bson:"_id,omitempty"`
	ProductID        bson.ObjectID  `bson:"product_id"`
	StoreID          bson.ObjectID  `bson:"store_id"`
	CustomerRoleID   *bson.ObjectID `bson:"customer_role_id"`
	Quantity         int            `bson:"quantity"`
	Price            float64        `bson:"price"`
	StartDateTimeUtc *time.Time     `bson:"start_date_time_utc"`
	EndDateTimeUtc   *time.Time     `bson:"end_date_time_utc"`
}

type TierPriceRepository interface {
	CreateMany(c context.Context, items []TierPrice) error
	Create(c context.Context, tier_price *TierPrice) error
	Update(c context.Context, tier_price *TierPrice) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]TierPrice, error)
	FetchByID(c context.Context, ID string) (TierPrice, error)
}

type TierPriceUsecase interface {
	CreateMany(c context.Context, items []TierPrice) error
	FetchByID(c context.Context, ID string) (TierPrice, error)
	Create(c context.Context, tier_price *TierPrice) error
	Update(c context.Context, tier_price *TierPrice) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]TierPrice, error)
}
