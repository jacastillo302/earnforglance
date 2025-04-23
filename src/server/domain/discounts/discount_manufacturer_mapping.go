package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDiscountManufacturerMapping = "discount_manufacturer_mappings"
)

// DiscountManufacturerMapping represents a discount-manufacturer mapping class
type DiscountManufacturerMapping struct {
	DiscountMapping
	EntityID bson.ObjectID `bson:"entity_id"`
}

type DiscountManufacturerMappingRepository interface {
	CreateMany(c context.Context, items []DiscountManufacturerMapping) error
	Create(c context.Context, discount_manufacturer_mapping *DiscountManufacturerMapping) error
	Update(c context.Context, discount_manufacturer_mapping *DiscountManufacturerMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountManufacturerMapping, error)
	FetchByID(c context.Context, discount_manufacturer_mappingID string) (DiscountManufacturerMapping, error)
}

type DiscountManufacturerMappingUsecase interface {
	CreateMany(c context.Context, items []DiscountManufacturerMapping) error
	FetchByID(c context.Context, discount_manufacturer_mappingID string) (DiscountManufacturerMapping, error)
	Create(c context.Context, discount_manufacturer_mapping *DiscountManufacturerMapping) error
	Update(c context.Context, discount_manufacturer_mapping *DiscountManufacturerMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DiscountManufacturerMapping, error)
}
