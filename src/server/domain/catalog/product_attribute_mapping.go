package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductAttributeMapping = "product_attribute_mappings"
)

// ProductAttributeMapping represents a product attribute mapping
type ProductAttributeMapping struct {
	ID                              bson.ObjectID `bson:"_id,omitempty"`
	ProductID                       bson.ObjectID `bson:"product_id"`
	ProductAttributeID              bson.ObjectID `bson:"product_attribute_id"`
	TextPrompt                      string        `bson:"text_prompt"`
	IsRequired                      bool          `bson:"is_required"`
	AttributeControlTypeID          int           `bson:"attribute_control_type_id"`
	DisplayOrder                    int           `bson:"display_order"`
	ValidationMinLength             *int          `bson:"validation_min_length"`
	ValidationMaxLength             *int          `bson:"validation_max_length"`
	ValidationFileAllowedExtensions string        `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int          `bson:"validation_file_maximum_size"`
	DefaultValue                    string        `bson:"default_value"`
	ConditionAttributeXml           string        `bson:"condition_attribute_xml"`
}

type ProductAttributeMappingRepository interface {
	CreateMany(c context.Context, items []ProductAttributeMapping) error
	Create(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Update(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeMapping, error)
	FetchByID(c context.Context, ID string) (ProductAttributeMapping, error)
}

type ProductAttributeMappingUsecase interface {
	CreateMany(c context.Context, items []ProductAttributeMapping) error
	FetchByID(c context.Context, ID string) (ProductAttributeMapping, error)
	Create(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Update(c context.Context, product_attribute_mapping *ProductAttributeMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeMapping, error)
}
