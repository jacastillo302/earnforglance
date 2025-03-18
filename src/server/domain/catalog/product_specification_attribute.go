package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductSpecificationAttribute = "product_specification_attributes"
)

// ProductSpecificationAttribute represents a product specification attribute
type ProductSpecificationAttribute struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	ProductID                      primitive.ObjectID `bson:"product_id"`
	AttributeTypeID                primitive.ObjectID `bson:"attribute_type_id"`
	SpecificationAttributeOptionID primitive.ObjectID `bson:"specification_attribute_option_id"`
	CustomValue                    string             `bson:"custom_value"`
	AllowFiltering                 bool               `bson:"allow_filtering"`
	ShowOnProductPage              bool               `bson:"show_on_product_page"`
	DisplayOrder                   int                `bson:"display_order"`
	AttributeType                  int                `bson:"attribute_type"`
}

type ProductSpecificationAttributeRepository interface {
	Create(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Update(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Delete(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Fetch(c context.Context) ([]ProductSpecificationAttribute, error)
	FetchByID(c context.Context, ProductSpecificationAttributeID string) (ProductSpecificationAttribute, error)
}

type ProductSpecificationAttributeUsecase interface {
	FetchByID(c context.Context, product_specification_attributeID string) (ProductSpecificationAttribute, error)
	Create(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Update(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Delete(c context.Context, product_specification_attribute *ProductSpecificationAttribute) error
	Fetch(c context.Context) ([]ProductSpecificationAttribute, error)
}
