package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPredefinedProductAttributeValue = "predefined_product_attribute_values"
)

// PredefinedProductAttributeValue represents a predefined (default) product attribute value
type PredefinedProductAttributeValue struct {
	ID                           bson.ObjectID `bson:"_id,omitempty"`
	ProductAttributeID           bson.ObjectID `bson:"product_attribute_id"`
	Name                         string        `bson:"name"`
	PriceAdjustment              float64       `bson:"price_adjustment"`
	PriceAdjustmentUsePercentage bool          `bson:"price_adjustment_use_percentage"`
	WeightAdjustment             float64       `bson:"weight_adjustment"`
	Cost                         float64       `bson:"cost"`
	IsPreSelected                bool          `bson:"is_pre_selected"`
	DisplayOrder                 int           `bson:"display_order"`
}

type PredefinedProductAttributeValueRepository interface {
	CreateMany(c context.Context, items []PredefinedProductAttributeValue) error
	Create(c context.Context, predefined_product_attribute_value *PredefinedProductAttributeValue) error
	Update(c context.Context, predefined_product_attribute_value *PredefinedProductAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PredefinedProductAttributeValue, error)
	FetchByID(c context.Context, ID string) (PredefinedProductAttributeValue, error)
}

type PredefinedProductAttributeValueUsecase interface {
	CreateMany(c context.Context, items []PredefinedProductAttributeValue) error
	FetchByID(c context.Context, ID string) (PredefinedProductAttributeValue, error)
	Create(c context.Context, predefined_product_attribute_value *PredefinedProductAttributeValue) error
	Update(c context.Context, predefined_product_attribute_value *PredefinedProductAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PredefinedProductAttributeValue, error)
}
