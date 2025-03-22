package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShippingMethod = "shipping_methods"
)

// ShippingMethod represents a shipping method (used by offline shipping rate computation methods).
type ShippingMethod struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	DisplayOrder int                `bson:"display_order"`
}

// ShippingMethodRepository defines the repository interface for ShippingMethod
type ShippingMethodRepository interface {
	Create(c context.Context, shipping_method *ShippingMethod) error
	Update(c context.Context, shipping_method *ShippingMethod) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingMethod, error)
	FetchByID(c context.Context, ID string) (ShippingMethod, error)
}

// ShippingMethodUsecase defines the use case interface for ShippingMethod
type ShippingMethodUsecase interface {
	FetchByID(c context.Context, ID string) (ShippingMethod, error)
	Create(c context.Context, shipping_method *ShippingMethod) error
	Update(c context.Context, shipping_method *ShippingMethod) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingMethod, error)
}
