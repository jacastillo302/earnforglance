package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductManufacturer = "product_manufacturers"
)

// ProductManufacturer represents a product manufacturer mapping
type ProductManufacturer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ProductID         int                `bson:"product_id"`
	ManufacturerID    int                `bson:"manufacturer_id"`
	IsFeaturedProduct bool               `bson:"is_featured_product"`
	DisplayOrder      int                `bson:"display_order"`
}
