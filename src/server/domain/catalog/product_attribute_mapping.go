package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductAttributeMapping = "product_attribute_mappings"
)

// ProductAttributeMapping represents a product attribute mapping
type ProductAttributeMapping struct {
	ID                              primitive.ObjectID   `bson:"_id,omitempty"`
	ProductID                       primitive.ObjectID   `bson:"product_id"`
	ProductAttributeID              primitive.ObjectID   `bson:"product_attribute_id"`
	TextPrompt                      string               `bson:"text_prompt"`
	IsRequired                      bool                 `bson:"is_required"`
	AttributeControlTypeID          int                  `bson:"attribute_control_type_id"`
	DisplayOrder                    int                  `bson:"display_order"`
	ValidationMinLength             *int                 `bson:"validation_min_length,omitempty"`
	ValidationMaxLength             *int                 `bson:"validation_max_length,omitempty"`
	ValidationFileAllowedExtensions string               `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int                 `bson:"validation_file_maximum_size,omitempty"`
	DefaultValue                    string               `bson:"default_value"`
	ConditionAttributeXml           string               `bson:"condition_attribute_xml"`
	AttributeControlType            AttributeControlType `bson:"attribute_control_type"`
}

type ProductAttributeMappingRepository interface {
	Create(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Update(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeMapping, error)
	FetchByID(c context.Context, ID string) (ProductAttributeMapping, error)
}

type ProductAttributeMappingUsecase interface {
	FetchByID(c context.Context, ID string) (ProductAttributeMapping, error)
	Create(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Update(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeMapping, error)
}
