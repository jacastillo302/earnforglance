package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductAvailabilityRange = "product_availability_ranges"
)

// ProductAvailabilityRange represents a product availability range
type ProductAvailabilityRange struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
