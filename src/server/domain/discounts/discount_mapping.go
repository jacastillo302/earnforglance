package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionDiscountMapping = "discount_mappings"
)

// DiscountMapping represents an abstract discount mapping
type DiscountMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DiscountID int                `bson:"discount_id"`
	EntityID   int                `bson:"entity_id"`
}

// NewDiscountMapping creates a new DiscountMapping instance
func NewDiscountMapping(discountID int, entityID int) *DiscountMapping {
	return &DiscountMapping{
		DiscountID: discountID,
		EntityID:   entityID,
	}
}
