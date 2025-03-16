package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductTemplate = "product_templates"
)

// ProductTemplate represents a product template
type ProductTemplate struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name"`
	ViewPath            string             `bson:"view_path"`
	DisplayOrder        int                `bson:"display_order"`
	IgnoredProductTypes string             `bson:"ignored_product_types"`
}
