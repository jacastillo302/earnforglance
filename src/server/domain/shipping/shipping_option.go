package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShippingOption = "shipping_options"
)

// ShippingOption represents a shipping option.
type ShippingOption struct {
	ID                                      primitive.ObjectID `bson:"_id,omitempty"`
	ShippingRateComputationMethodSystemName string             `bson:"shipping_rate_computation_method_system_name"`
	Rate                                    float64            `bson:"rate"`
	Name                                    string             `bson:"name"`
	Description                             string             `bson:"description"`
	TransitDays                             *int               `bson:"transit_days,omitempty"`
	IsPickupInStore                         bool               `bson:"is_pickup_in_store"`
	DisplayOrder                            *int               `bson:"display_order,omitempty"`
}

// ShippingOptionRepository defines the repository interface for ShippingOption
type ShippingOptionRepository interface {
	CreateMany(c context.Context, items []ShippingOption) error
	Create(c context.Context, shipping_option *ShippingOption) error
	Update(c context.Context, shipping_option *ShippingOption) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingOption, error)
	FetchByID(c context.Context, ID string) (ShippingOption, error)
}

// ShippingOptionUsecase defines the usecase interface for ShippingOption
type ShippingOptionUsecase interface {
	FetchByID(c context.Context, ID string) (ShippingOption, error)
	Create(c context.Context, shipping_option *ShippingOption) error
	Update(c context.Context, shipping_option *ShippingOption) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingOption, error)
}
