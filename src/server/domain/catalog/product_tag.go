package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductTag = "product_tags"
)

// ProductTag represents a product tag
type ProductTag struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	MetaDescription string             `bson:"meta_description"`
	MetaKeywords    string             `bson:"meta_keywords"`
	MetaTitle       string             `bson:"meta_title"`
}
