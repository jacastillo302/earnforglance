package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionAddressAttributeValue = "address_attribute_values"
)

// AddressAttributeValue represents an address attribute value
type AddressAttributeValue struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	AddressAttributeID int                `bson:"address_attribute_id"`
	Name               string             `bson:"name"`
	IsPreSelected      bool               `bson:"is_pre_selected"`
	DisplayOrder       int                `bson:"display_order"`
}
