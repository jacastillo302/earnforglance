package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductProductTagMapping = "product_product_tag_mappings"
)

// ProductProductTagMapping represents a product-product tag mapping class
type ProductProductTagMapping struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	ProductID    bson.ObjectID `bson:"product_id"`
	ProductTagID bson.ObjectID `bson:"product_tag_id"`
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
