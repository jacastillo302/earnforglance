package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductManufacturer = "product_manufacturers"
)

// ProductManufacturer represents a product manufacturer mapping
type ProductManufacturer struct {
	ID                bson.ObjectID `bson:"_id,omitempty"`
	ProductID         bson.ObjectID `bson:"product_id"`
	ManufacturerID    bson.ObjectID `bson:"manufacturer_id"`
	IsFeaturedProduct bool          `bson:"is_featured_product"`
	DisplayOrder      int           `bson:"display_order"`
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
	CreateMany(c context.Context, items []ProductManufacturer) error
	FetchByID(c context.Context, ID string) (ProductManufacturer, error)
	Create(c context.Context, product_manufacturer *ProductManufacturer) error
	Update(c context.Context, product_manufacturer *ProductManufacturer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductManufacturer, error)
}
