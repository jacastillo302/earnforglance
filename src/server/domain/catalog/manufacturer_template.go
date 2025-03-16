package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionManufacturerTemplate = "manufacturer_templates"
)

// ManufacturerTemplate represents a manufacturer template
type ManufacturerTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}
