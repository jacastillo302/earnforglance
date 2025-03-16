package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCategoryTemplate = "category_templates"
)

// CategoryTemplate represents a category template
type CategoryTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}
