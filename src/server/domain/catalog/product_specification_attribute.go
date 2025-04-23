package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductSpecificationAttribute = "product_specification_attributes"
)

// ProductSpecificationAttribute represents a product specification attribute
type ProductSpecificationAttribute struct {
	ProductID                      bson.ObjectID `bson:"product_id"`
	SpecificationAttributeTypeID   int           `bson:"attribute_type_id"`
	SpecificationAttributeOptionID bson.ObjectID `bson:"specification_attribute_option_id"`
	CustomValue                    string        `bson:"custom_value"`
	AllowFiltering                 bool          `bson:"allow_filtering"`
	ShowOnProductPage              bool          `bson:"show_on_product_page"`
	DisplayOrder                   int           `bson:"display_order"`
}

type ProductSpecificationAttributeRepository interface {
	CreateMany(c context.Context, items []ProductSpecificationAttribute) error
	Create(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Update(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductSpecificationAttribute, error)
	FetchByID(c context.Context, ID string) (ProductSpecificationAttribute, error)
}

type ProductSpecificationAttributeUsecase interface {
	CreateMany(c context.Context, items []ProductSpecificationAttribute) error
	FetchByID(c context.Context, ID string) (ProductSpecificationAttribute, error)
	Create(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Update(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductSpecificationAttribute, error)
}
