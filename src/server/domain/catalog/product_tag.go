package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductTag = "product_tags"
)

// ProductTag represents a product tag
type ProductTag struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Name            string        `bson:"name"`
	MetaDescription string        `bson:"meta_description"`
	MetaKeywords    string        `bson:"meta_keywords"`
	MetaTitle       string        `bson:"meta_title"`
}

type ProductTagRepository interface {
	CreateMany(c context.Context, items []ProductTag) error
	Create(c context.Context, product_tag *ProductTag) error
	Update(c context.Context, product_tag *ProductTag) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductTag, error)
	FetchByID(c context.Context, ID string) (ProductTag, error)
}

type ProductTagUsecase interface {
	CreateMany(c context.Context, items []ProductTag) error
	FetchByID(c context.Context, ID string) (ProductTag, error)
	Create(c context.Context, product_tag *ProductTag) error
	Update(c context.Context, product_tag *ProductTag) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductTag, error)
}
