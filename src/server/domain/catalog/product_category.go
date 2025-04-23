package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductCategory = "product_categories"
)

// ProductCategory represents a product category mapping
type ProductCategory struct {
	ID                bson.ObjectID `bson:"_id,omitempty"`
	ProductID         bson.ObjectID `bson:"product_id"`
	CategoryID        bson.ObjectID `bson:"category_id"`
	IsFeaturedProduct bool          `bson:"is_featured_product"`
	DisplayOrder      int           `bson:"display_order"`
}

type ProductCategoryRepository interface {
	CreateMany(c context.Context, items []ProductCategory) error
	Create(c context.Context, product_category *ProductCategory) error
	Update(c context.Context, product_category *ProductCategory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductCategory, error)
	FetchByID(c context.Context, ID string) (ProductCategory, error)
}

type ProductCategoryUsecase interface {
	CreateMany(c context.Context, items []ProductCategory) error
	FetchByID(c context.Context, ID string) (ProductCategory, error)
	Create(c context.Context, product_category *ProductCategory) error
	Update(c context.Context, product_category *ProductCategory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductCategory, error)
}
