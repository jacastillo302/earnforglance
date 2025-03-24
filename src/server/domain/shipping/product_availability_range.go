package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductAvailabilityRange = "product_availability_ranges"
)

// ProductAvailabilityRange represents a product availability range.
type ProductAvailabilityRange struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}

// ProductAvailabilityRangeRepository defines the repository interface for ProductAvailabilityRange
type ProductAvailabilityRangeRepository interface {
	CreateMany(c context.Context, items []ProductAvailabilityRange) error
	Create(c context.Context, product_availability_range *ProductAvailabilityRange) error
	Update(c context.Context, product_availability_range *ProductAvailabilityRange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAvailabilityRange, error)
	FetchByID(c context.Context, ID string) (ProductAvailabilityRange, error)
}

// ProductAvailabilityRangeUsecase defines the use case interface for ProductAvailabilityRange
type ProductAvailabilityRangeUsecase interface {
	FetchByID(c context.Context, ID string) (ProductAvailabilityRange, error)
	Create(c context.Context, product_availability_range *ProductAvailabilityRange) error
	Update(c context.Context, product_availability_range *ProductAvailabilityRange) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAvailabilityRange, error)
}
