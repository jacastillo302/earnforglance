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
	ProductTagID int                `bson:"product_tag_id"`
}

type ProductProductTagMappingRepository interface {
	Create(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Update(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Delete(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Fetch(c context.Context) ([]ProductProductTagMapping, error)
	FetchByID(c context.Context, product_product_tag_mappingsID string) (ProductProductTagMapping, error)
}

type ProductProductTagMappingUsecase interface {
	FetchByID(c context.Context, product_product_tag_mappingsID string) (ProductProductTagMapping, error)
	Create(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Update(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Delete(c context.Context, product_product_tag_mappings *ProductProductTagMapping) error
	Fetch(c context.Context) ([]ProductProductTagMapping, error)
}
