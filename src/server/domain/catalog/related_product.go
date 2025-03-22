package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRelatedProduct = "related_products"
)

// RelatedProduct represents a related product
type RelatedProduct struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID1   primitive.ObjectID `bson:"product_id1"`
	ProductID2   primitive.ObjectID `bson:"product_id2"`
	DisplayOrder int                `bson:"display_order"`
}

type RelatedProductRepository interface {
	Create(c context.Context, related_product *RelatedProduct) error
	Update(c context.Context, related_product *RelatedProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RelatedProduct, error)
	FetchByID(c context.Context, ID string) (RelatedProduct, error)
}

type RelatedProductUsecase interface {
	FetchByID(c context.Context, ID string) (RelatedProduct, error)
	Create(c context.Context, related_product *RelatedProduct) error
	Update(c context.Context, related_product *RelatedProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RelatedProduct, error)
}
