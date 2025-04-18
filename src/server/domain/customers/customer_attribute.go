package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerAttribute = "customer_customer_attributes"
)

// CustomerAttribute represents a customer attribute
type CustomerAttribute struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty"`
	Name                            string             `bson:"name"`
	IsRequired                      bool               `bson:"is_required"`
	AttributeControlTypeID          int                `bson:"attribute_control_type_id"`
	DisplayOrder                    int                `bson:"display_order"`
	DefaultValue                    string             `bson:"default_value"`
	ValidationMinLength             *int               `bson:"validation_min_length"`
	ValidationMaxLength             *int               `bson:"validation_max_length"`
	ValidationFileAllowedExtensions string             `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int               `bson:"validation_file_maximum_size"`
	ConditionAttributeXml           string             `bson:"condition_attribute_xml"`
}

type CustomerAttributeRepository interface {
	CreateMany(c context.Context, items []CustomerAttribute) error
	Create(c context.Context, customer_attribute *CustomerAttribute) error
	Update(c context.Context, customer_attribute *CustomerAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAttribute, error)
	FetchByID(c context.Context, ID string) (CustomerAttribute, error)
}

type CustomerAttributeUsecase interface {
	CreateMany(c context.Context, items []CustomerAttribute) error
	FetchByID(c context.Context, ID string) (CustomerAttribute, error)
	Create(c context.Context, customer_attribute *CustomerAttribute) error
	Update(c context.Context, customer_attribute *CustomerAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAttribute, error)
}
