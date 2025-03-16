package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCheckoutAttributeValue = "checkout_attribute_values"
)

// CheckoutAttributeValue represents a checkout attribute value
type CheckoutAttributeValue struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ColorSquaresRgb  string             `bson:"color_squares_rgb"`
	PriceAdjustment  float64            `bson:"price_adjustment"`
	WeightAdjustment float64            `bson:"weight_adjustment"`
}
