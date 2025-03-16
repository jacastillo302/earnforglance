package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionBaseAttributeValue = "base_attribute_values"
)

// BaseAttributeValue represents the base class for attribute values
type BaseAttributeValue struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	IsPreSelected bool               `bson:"is_pre_selected,omitempty"`
	DisplayOrder  int                `bson:"display_order,omitempty"`
	AttributeId   int                `bson:"attribute_id,omitempty"`
}
