package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShippingMethodCountryMapping = "shipping_method_country_mappings"
)

// ShippingMethodCountryMapping represents a shipping method-country mapping class
type ShippingMethodCountryMapping struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ShippingMethodID primitive.ObjectID `bson:"shipping_method_id"`
	CountryID        primitive.ObjectID `bson:"country_id"`
}

// ShippingMethodCountryMappingRepository defines the repository interface for ShippingMethodCountryMapping
type ShippingMethodCountryMappingRepository interface {
	CreateMany(c context.Context, items []ShippingMethodCountryMapping) error
	Create(c context.Context, shipping_method_country *ShippingMethodCountryMapping) error
	Update(c context.Context, shipping_method_country *ShippingMethodCountryMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingMethodCountryMapping, error)
	FetchByID(c context.Context, ID string) (ShippingMethodCountryMapping, error)
}

// ShippingMethodCountryMappingUsecase defines the use case interface for ShippingMethodCountryMapping
type ShippingMethodCountryMappingUsecase interface {
	FetchByID(c context.Context, ID string) (ShippingMethodCountryMapping, error)
	Create(c context.Context, shipping_method_country *ShippingMethodCountryMapping) error
	Update(c context.Context, shipping_method_country *ShippingMethodCountryMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingMethodCountryMapping, error)
}
