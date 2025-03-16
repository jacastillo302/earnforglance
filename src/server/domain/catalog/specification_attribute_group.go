package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSpecificationAttributeGroup = "specification_attribute_groups"
)

// SpecificationAttributeGroup represents a specification attribute group
type SpecificationAttributeGroup struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
