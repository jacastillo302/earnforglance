package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDiscountProductMapping = "discount_product_mappings"
)

// DiscountProductMapping represents a discount-product mapping class
type DiscountProductMapping struct {
	DiscountMapping
	EntityID bson.ObjectID `bson:"entity_id"`
}

type DiscountProductMappingRepository interface {
	CreateMany(c context.Context, items []DiscountProductMapping) error
	Create(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Update(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountProductMapping, error)
	FetchByID(c context.Context, ID string) (DiscountProductMapping, error)
}

type DiscountProductMappingUsecase interface {
	CreateMany(c context.Context, items []DiscountProductMapping) error
	FetchByID(c context.Context, ID string) (DiscountProductMapping, error)
	Create(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Update(c context.Context, discount_product_mapping *DiscountProductMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountProductMapping, error)
}
