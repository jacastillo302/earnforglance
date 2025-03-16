package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionRelatedProduct = "related_products"
)

// RelatedProduct represents a related product
type RelatedProduct struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID1   int                `bson:"product_id1"`
	ProductID2   int                `bson:"product_id2"`
	DisplayOrder int                `bson:"display_order"`
}
