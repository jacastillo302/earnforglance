package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPredefinedProductAttributeValue = "predefined_product_attribute_values"
)

// PredefinedProductAttributeValue represents a predefined (default) product attribute value
type PredefinedProductAttributeValue struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty"`
	ProductAttributeID           int                `bson:"product_attribute_id"`
	Name                         string             `bson:"name"`
	PriceAdjustment              float64            `bson:"price_adjustment"`
	PriceAdjustmentUsePercentage bool               `bson:"price_adjustment_use_percentage"`
	WeightAdjustment             float64            `bson:"weight_adjustment"`
	Cost                         float64            `bson:"cost"`
	IsPreSelected                bool               `bson:"is_pre_selected"`
	DisplayOrder                 int                `bson:"display_order"`
}
