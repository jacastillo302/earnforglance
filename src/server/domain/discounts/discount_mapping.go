package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountMapping = "discount_mappings"
)

// DiscountMapping represents an abstract discount mapping
type DiscountMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	DiscountID primitive.ObjectID `bson:"discount_id"`
	EntityID   primitive.ObjectID `bson:"entity_id"`
}

type DiscountMappingRepository interface {
	CreateMany(c context.Context, items []DiscountMapping) error
	Create(c context.Context, discount_mapping *DiscountMapping) error
	Update(c context.Context, discount_mapping *DiscountMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountMapping, error)
	FetchByID(c context.Context, ID string) (DiscountMapping, error)
}

type DiscountMappingUsecase interface {
	FetchByID(c context.Context, ID string) (DiscountMapping, error)
	Create(c context.Context, discount_mapping *DiscountMapping) error
	Update(c context.Context, discount_mapping *DiscountMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountMapping, error)
}

// NewDiscountMapping creates a new DiscountMapping instance
func NewDiscountMapping(discountID primitive.ObjectID, entityID primitive.ObjectID) *DiscountMapping {
	return &DiscountMapping{
		DiscountID: discountID,
		EntityID:   entityID,
	}
}
