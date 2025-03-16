package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductCategory = "product_categories"
)

// ProductCategory represents a product category mapping
type ProductCategory struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ProductID         int                `bson:"product_id"`
	CategoryID        int                `bson:"category_id"`
	IsFeaturedProduct bool               `bson:"is_featured_product"`
	DisplayOrder      int                `bson:"display_order"`
}
