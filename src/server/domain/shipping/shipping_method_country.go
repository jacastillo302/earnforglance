package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionShippingMethodCountryMapping = "shipping_method_country_mappings"
)

// ShippingMethodCountryMapping represents a shipping method-country mapping class
type ShippingMethodCountryMapping struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ShippingMethodID int                `bson:"shipping_method_id"`
	CountryID        int                `bson:"country_id"`
}
