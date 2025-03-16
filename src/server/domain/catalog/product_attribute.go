package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductAttribute = "product_attributes"
)

// ProductAttribute represents a product attribute
type ProductAttribute struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}
