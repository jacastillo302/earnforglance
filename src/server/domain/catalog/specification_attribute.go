package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSpecificationAttribute = "specification_attributes"
)

// SpecificationAttribute represents a specification attribute
type SpecificationAttribute struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`
	Name                          string             `bson:"name"`
	DisplayOrder                  int                `bson:"display_order"`
	SpecificationAttributeGroupID *int               `bson:"specification_attribute_group_id,omitempty"`
}
