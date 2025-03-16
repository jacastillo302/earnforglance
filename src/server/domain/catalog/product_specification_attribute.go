package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductSpecificationAttribute = "product_specification_attributes"
)

// ProductSpecificationAttribute represents a product specification attribute
type ProductSpecificationAttribute struct {
	ID                             primitive.ObjectID         `bson:"_id,omitempty"`
	ProductID                      int                        `bson:"product_id"`
	AttributeTypeID                int                        `bson:"attribute_type_id"`
	SpecificationAttributeOptionID int                        `bson:"specification_attribute_option_id"`
	CustomValue                    string                     `bson:"custom_value"`
	AllowFiltering                 bool                       `bson:"allow_filtering"`
	ShowOnProductPage              bool                       `bson:"show_on_product_page"`
	DisplayOrder                   int                        `bson:"display_order"`
	AttributeType                  SpecificationAttributeType `bson:"attribute_type"`
}
