package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionRelatedProduct = "related_products"
)

// RelatedProduct represents a related product
type RelatedProduct struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	ProductID1   bson.ObjectID `bson:"product_id1"`
	ProductID2   bson.ObjectID `bson:"product_id2"`
	DisplayOrder int           `bson:"display_order"`
}

type RelatedProductRepository interface {
	CreateMany(c context.Context, items []RelatedProduct) error
	Create(c context.Context, related_product *RelatedProduct) error
	Update(c context.Context, related_product *RelatedProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RelatedProduct, error)
	FetchByID(c context.Context, ID string) (RelatedProduct, error)
}

type RelatedProductUsecase interface {
	CreateMany(c context.Context, items []RelatedProduct) error
	FetchByID(c context.Context, ID string) (RelatedProduct, error)
	Create(c context.Context, related_product *RelatedProduct) error
	Update(c context.Context, related_product *RelatedProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RelatedProduct, error)
}
