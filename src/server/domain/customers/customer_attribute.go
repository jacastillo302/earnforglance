package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCustomerCustomerAttribute = "customer_customer_attributes"
)

// CustomerAttribute represents a customer attribute
type CustomerAttribute struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty"`
	Name                            string             `bson:"name"`
	IsRequired                      bool               `bson:"is_required"`
	AttributeControlTypeID          int                `bson:"attribute_control_type_id"`
	DisplayOrder                    int                `bson:"display_order"`
	DefaultValue                    string             `bson:"default_value"`
	ValidationMinLength             *int               `bson:"validation_min_length,omitempty"`
	ValidationMaxLength             *int               `bson:"validation_max_length,omitempty"`
	ValidationFileAllowedExtensions string             `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int               `bson:"validation_file_maximum_size,omitempty"`
	ConditionAttributeXml           string             `bson:"condition_attribute_xml"`
}
