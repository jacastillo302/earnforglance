package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductProductTagMapping = "product_product_tag_mappings"
)

// ProductProductTagMapping represents a product-product tag mapping class
type ProductProductTagMapping struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    primitive.ObjectID `bson:"product_id"`
	ProductTagID primitive.ObjectID `bson:"product_tag_id"`
}

type ProductProductTagMappingRepository interface {
	CreateMany(c context.Context, items []ProductProductTagMapping) error
	Create(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Update(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductProductTagMapping, error)
	FetchByID(c context.Context, ID string) (ProductProductTagMapping, error)
}

type ProductProductTagMappingUsecase interface {
	CreateMany(c context.Context, items []ProductProductTagMapping) error
	FetchByID(c context.Context, ID string) (ProductProductTagMapping, error)
	Create(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Update(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductProductTagMapping, error)
}
