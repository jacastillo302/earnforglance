package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductTag = "product_tags"
)

// ProductTag represents a product tag
type ProductTag struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	MetaDescription string             `bson:"meta_description"`
	MetaKeywords    string             `bson:"meta_keywords"`
	MetaTitle       string             `bson:"meta_title"`
}

type ProductTagRepository interface {
	Create(c context.Context, product_tag *ProductTag) error
	Update(c context.Context, product_tag *ProductTag) error
	Delete(c context.Context, product_tag *ProductTag) error
	Fetch(c context.Context) ([]ProductTag, error)
	FetchByID(c context.Context, product_tagID string) (ProductTag, error)
}

type ProductTagUsecase interface {
	FetchByID(c context.Context, product_tagID string) (ProductTag, error)
	Create(c context.Context, product_tag *ProductTag) error
	Update(c context.Context, product_tag *ProductTag) error
	Delete(c context.Context, product_tag *ProductTag) error
	Fetch(c context.Context) ([]ProductTag, error)
}
