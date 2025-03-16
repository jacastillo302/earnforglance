package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionShippingMethod = "shipping_methods"
)

// ShippingMethod represents a shipping method (used by offline shipping rate computation methods)
type ShippingMethod struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	DisplayOrder int                `bson:"display_order"`
}
