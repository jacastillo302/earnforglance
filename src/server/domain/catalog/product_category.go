package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductCategory = "product_categories"
)

// ProductCategory represents a product category mapping
type ProductCategory struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ProductID         primitive.ObjectID `bson:"product_id"`
	CategoryID        primitive.ObjectID `bson:"category_id"`
	IsFeaturedProduct bool               `bson:"is_featured_product"`
	DisplayOrder      int                `bson:"display_order"`
}

type ProductCategoryRepository interface {
	Create(c context.Context, product_category *ProductCategory) error
	Update(c context.Context, product_category *ProductCategory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductCategory, error)
	FetchByID(c context.Context, ID string) (ProductCategory, error)
}

type ProductCategoryUsecase interface {
	FetchByID(c context.Context, ID string) (ProductCategory, error)
	Create(c context.Context, product_category *ProductCategory) error
	Update(c context.Context, product_category *ProductCategory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductCategory, error)
}
