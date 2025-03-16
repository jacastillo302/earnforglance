package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSpecificationAttributeOption = "specification_attribute_options"
)

// SpecificationAttributeOption represents a specification attribute option
type SpecificationAttributeOption struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	SpecificationAttributeID int                `bson:"specification_attribute_id"`
	Name                     string             `bson:"name"`
	ColorSquaresRgb          string             `bson:"color_squares_rgb"`
	DisplayOrder             int                `bson:"display_order"`
}
