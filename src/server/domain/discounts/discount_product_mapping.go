package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDiscountProductMapping = "discount_product_mappings"
)

// DiscountProductMapping represents a discount-product mapping class
type DiscountProductMapping struct {
	DiscountMapping
	EntityID primitive.ObjectID `bson:"entity_id"`
}

type DiscountProductMappingRepository interface {
	Create(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Update(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountProductMapping, error)
	FetchByID(c context.Context, ID string) (DiscountProductMapping, error)
}

type DiscountProductMappingUsecase interface {
	FetchByID(c context.Context, ID string) (DiscountProductMapping, error)
	Create(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Update(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountProductMapping, error)
}
