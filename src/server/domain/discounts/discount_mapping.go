package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDiscountMapping = "discount_mappings"
)

// DiscountMapping represents an abstract discount mapping
type DiscountMapping struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	DiscountID bson.ObjectID `bson:"discount_id"`
	EntityID   bson.ObjectID `bson:"entity_id"`
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
	CreateMany(c context.Context, items []DiscountMapping) error
	FetchByID(c context.Context, ID string) (DiscountMapping, error)
	Create(c context.Context, discount_mapping *DiscountMapping) error
	Update(c context.Context, discount_mapping *DiscountMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountMapping, error)
}

// NewDiscountMapping creates a new DiscountMapping instance
func NewDiscountMapping(discountID bson.ObjectID, entityID bson.ObjectID) *DiscountMapping {
	return &DiscountMapping{
		DiscountID: discountID,
		EntityID:   entityID,
	}
}
