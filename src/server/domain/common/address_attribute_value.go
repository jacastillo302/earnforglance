package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAddressAttributeValue = "address_attribute_values"
)

// AddressAttributeValue represents an address attribute value
type AddressAttributeValue struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	AddressAttributeID primitive.ObjectID `bson:"address_attribute_id"`
	Name               string             `bson:"name"`
	IsPreSelected      bool               `bson:"is_pre_selected"`
	DisplayOrder       int                `bson:"display_order"`
}

type AddressAttributeValueRepository interface {
	Create(c context.Context, address_attribute_value *AddressAttributeValue) error
	Update(c context.Context, address_attribute_value *AddressAttributeValue) error
	Delete(c context.Context, address_attribute_value *AddressAttributeValue) error
	Fetch(c context.Context) ([]AddressAttributeValue, error)
	FetchByID(c context.Context, address_attribute_valueID string) (AddressAttributeValue, error)
}

type AddressAttributeValueUsecase interface {
	FetchByID(c context.Context, address_attribute_valueID string) (AddressAttributeValue, error)
	Create(c context.Context, address_attribute_value *AddressAttributeValue) error
	Update(c context.Context, address_attribute_value *AddressAttributeValue) error
	Delete(c context.Context, address_attribute_value *AddressAttributeValue) error
	Fetch(c context.Context) ([]AddressAttributeValue, error)
}
