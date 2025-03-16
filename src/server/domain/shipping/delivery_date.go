package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionDeliveryDate = "delivery_dates"
)

// DeliveryDate represents a delivery date
type DeliveryDate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
