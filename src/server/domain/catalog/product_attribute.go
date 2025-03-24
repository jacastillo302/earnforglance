package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductAttribute = "product_attributes"
)

// ProductAttribute represents a product attribute
type ProductAttribute struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

type ProductAttributeRepository interface {
	CreateMany(c context.Context, items []ProductAttribute) error
	Create(c context.Context, product_attribute *ProductAttribute) error
	Update(c context.Context, product_attribute *ProductAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttribute, error)
	FetchByID(c context.Context, ID string) (ProductAttribute, error)
}

type ProductAttributeUsecase interface {
	CreateMany(c context.Context, items []ProductAttribute) error
	FetchByID(c context.Context, ID string) (ProductAttribute, error)
	Create(c context.Context, product_attribute *ProductAttribute) error
	Update(c context.Context, product_attribute *ProductAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttribute, error)
}
