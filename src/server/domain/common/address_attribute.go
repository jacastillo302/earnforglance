package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAddressAttribute = "address_attribute"
)

// AddressAttribute represents an address attribute
type AddressAttribute struct {
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

type AddressAttributeRepository interface {
	CreateMany(c context.Context, items []AddressAttribute) error
	Create(c context.Context, address_attribute *AddressAttribute) error
	Update(c context.Context, address_attribute *AddressAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AddressAttribute, error)
	FetchByID(c context.Context, ID string) (AddressAttribute, error)
}

type AddressAttributeUsecase interface {
	FetchByID(c context.Context, ID string) (AddressAttribute, error)
	Create(c context.Context, address_attribute *AddressAttribute) error
	Update(c context.Context, address_attribute *AddressAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AddressAttribute, error)
}
