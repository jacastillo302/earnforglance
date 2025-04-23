package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductTemplate = "product_templates"
)

// ProductTemplate represents a product template
type ProductTemplate struct {
	ID                  bson.ObjectID `bson:"_id,omitempty"`
	Name                string        `bson:"name"`
	ViewPath            string        `bson:"view_path"`
	DisplayOrder        int           `bson:"display_order"`
	IgnoredProductTypes string        `bson:"ignored_product_types"`
}

type ProductTemplateRepository interface {
	CreateMany(c context.Context, items []ProductTemplate) error
	Create(c context.Context, product_template *ProductTemplate) error
	Update(c context.Context, product_template *ProductTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductTemplate, error)
	FetchByID(c context.Context, ID string) (ProductTemplate, error)
}

type ProductTemplateUsecase interface {
	CreateMany(c context.Context, items []ProductTemplate) error
	FetchByID(c context.Context, ID string) (ProductTemplate, error)
	Create(c context.Context, product_template *ProductTemplate) error
	Update(c context.Context, product_template *ProductTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductTemplate, error)
}
