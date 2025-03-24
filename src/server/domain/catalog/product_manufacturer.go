package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductManufacturer = "product_manufacturers"
)

// ProductManufacturer represents a product manufacturer mapping
type ProductManufacturer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ProductID         primitive.ObjectID `bson:"product_id"`
	ManufacturerID    primitive.ObjectID `bson:"manufacturer_id"`
	IsFeaturedProduct bool               `bson:"is_featured_product"`
	DisplayOrder      int                `bson:"display_order"`
}

type ProductManufacturerRepository interface {
	CreateMany(c context.Context, items []ProductManufacturer) error
	Create(c context.Context, product_manufacturer *ProductManufacturer) error
	Update(c context.Context, product_manufacturer *ProductManufacturer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductManufacturer, error)
	FetchByID(c context.Context, ID string) (ProductManufacturer, error)
}

type ProductManufacturerUsecase interface {
	FetchByID(c context.Context, ID string) (ProductManufacturer, error)
	Create(c context.Context, product_manufacturer *ProductManufacturer) error
	Update(c context.Context, product_manufacturer *ProductManufacturer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductManufacturer, error)
}
