package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductAttributeValue = "product_attribute_values"
)

// ProductAttributeValue represents a product attribute value
type ProductAttributeValue struct {
	ID                           bson.ObjectID  `bson:"_id,omitempty"`
	ProductAttributeMappingID    bson.ObjectID  `bson:"product_attribute_mapping_id"`
	AttributeValueTypeID         int            `bson:"attribute_value_type_id"`
	AssociatedProductID          bson.ObjectID  `bson:"associated_product_id"`
	Name                         string         `bson:"name"`
	ColorSquaresRgb              string         `bson:"color_squares_rgb"`
	ImageSquaresPictureID        bson.ObjectID  `bson:"image_squares_picture_id"`
	PriceAdjustment              float64        `bson:"price_adjustment"`
	PriceAdjustmentUsePercentage bool           `bson:"price_adjustment_use_percentage"`
	WeightAdjustment             float64        `bson:"weight_adjustment"`
	Cost                         float64        `bson:"cost"`
	CustomerEntersQty            bool           `bson:"customer_enters_qty"`
	Quantity                     int            `bson:"quantity"`
	IsPreSelected                bool           `bson:"is_pre_selected"`
	DisplayOrder                 int            `bson:"display_order"`
	PictureID                    *bson.ObjectID `bson:"picture_id"` // Deprecated field
}

type ProductAttributeValueRepository interface {
	CreateMany(c context.Context, items []ProductAttributeValue) error
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValue) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValue, error)
	FetchByID(c context.Context, ID string) (ProductAttributeValue, error)
}

type ProductAttributeValueUsecase interface {
	CreateMany(c context.Context, items []ProductAttributeValue) error
	FetchByID(c context.Context, ID string) (ProductAttributeValue, error)
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValue) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValue, error)
}
