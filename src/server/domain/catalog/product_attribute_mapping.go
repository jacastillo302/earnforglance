package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductAttributeMapping = "product_attribute_mappings"
)

// ProductAttributeMapping represents a product attribute mapping
type ProductAttributeMapping struct {
	ID                              primitive.ObjectID   `bson:"_id,omitempty"`
	ProductID                       int                  `bson:"product_id"`
	ProductAttributeID              int                  `bson:"product_attribute_id"`
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
