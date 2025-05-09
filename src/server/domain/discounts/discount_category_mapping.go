package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDiscountCategoryMapping = "discount_category_mappings"
)

// DiscountCategoryMapping represents a discount-category mapping class
type DiscountCategoryMapping struct {
	DiscountMapping
	EntityID bson.ObjectID `bson:"entity_id"`
}

type DiscountCategoryMappingRepository interface {
	CreateMany(c context.Context, items []DiscountCategoryMapping) error
	Create(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Update(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountCategoryMapping, error)
	FetchByID(c context.Context, ID string) (DiscountCategoryMapping, error)
}

type DiscountCategoryMappingUsecase interface {
	CreateMany(c context.Context, items []DiscountCategoryMapping) error
	FetchByID(c context.Context, ID string) (DiscountCategoryMapping, error)
	Create(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Update(c context.Context, discount_category_mapping *DiscountCategoryMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountCategoryMapping, error)
}
