package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCheckoutAttributeValue = "checkout_attribute_values"
)

// CheckoutAttributeValue represents a checkout attribute value.
type CheckoutAttributeValue struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	AttributeID      bson.ObjectID `bson:"attribute_id"`
	Name             string        `bson:"name"`
	DisplayOrder     int           `bson:"display_order"`
	IsPreSelected    bool          `bson:"is_pre_selected"`
	ColorSquaresRgb  string        `bson:"color_squares_rgb"`
	PriceAdjustment  float64       `bson:"price_adjustment"`
	WeightAdjustment float64       `bson:"weight_adjustment"`
}

// CheckoutAttributeValueRepository interface
type CheckoutAttributeValueRepository interface {
	CreateMany(c context.Context, items []CheckoutAttributeValue) error
	Create(c context.Context, checkout_attribute_value *CheckoutAttributeValue) error
	Update(c context.Context, checkout_attribute_value *CheckoutAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CheckoutAttributeValue, error)
	FetchByID(c context.Context, ID string) (CheckoutAttributeValue, error)
}

// CheckoutAttributeValueUsecase interface
type CheckoutAttributeValueUsecase interface {
	CreateMany(c context.Context, items []CheckoutAttributeValue) error
	FetchByID(c context.Context, ID string) (CheckoutAttributeValue, error)
	Create(c context.Context, checkout_attribute_value *CheckoutAttributeValue) error
	Update(c context.Context, checkout_attribute_value *CheckoutAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CheckoutAttributeValue, error)
}
