package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountCategoryMapping = "discount_category_mappings"
)

// DiscountCategoryMapping represents a discount-category mapping class
type DiscountCategoryMapping struct {
	DiscountMapping
	EntityID primitive.ObjectID `bson:"entity_id"`
}

type DiscountCategoryMappingRepository interface {
	Create(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Update(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Delete(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Fetch(c context.Context) ([]DiscountCategoryMapping, error)
	FetchByID(c context.Context, discount_category_mappingID string) (DiscountCategoryMapping, error)
}

type DiscountCategoryMappingUsecase interface {
	FetchByID(c context.Context, discount_category_mappingID string) (DiscountCategoryMapping, error)
	Create(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Update(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Delete(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Fetch(c context.Context) ([]DiscountCategoryMapping, error)
}
