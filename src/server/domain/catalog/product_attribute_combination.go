package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductAttributeCombination = "product_attribute_combinations"
)

// ProductAttributeCombination represents a product attribute combination
type ProductAttributeCombination struct {
	ID                          bson.ObjectID `bson:"_id,omitempty"`
	ProductID                   bson.ObjectID `bson:"product_id"`
	AttributesXml               string        `bson:"attributes_xml"`
	StockQuantity               int           `bson:"stock_quantity"`
	AllowOutOfStockOrders       bool          `bson:"allow_out_of_stock_orders"`
	Sku                         string        `bson:"sku"`
	ManufacturerPartNumber      string        `bson:"manufacturer_part_number"`
	Gtin                        string        `bson:"gtin"`
	OverriddenPrice             *float64      `bson:"overridden_price"`
	NotifyAdminForQuantityBelow int           `bson:"notify_admin_for_quantity_below"`
	MinStockQuantity            int           `bson:"min_stock_quantity"`
}

type ProductAttributeCombinationRepository interface {
	CreateMany(c context.Context, items []ProductAttributeCombination) error
	Create(c context.Context, product_attribute_combination *ProductAttributeCombination) error
	Update(c context.Context, product_attribute_combination *ProductAttributeCombination) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeCombination, error)
	FetchByID(c context.Context, ID string) (ProductAttributeCombination, error)
}

type ProductAttributeCombinationUsecase interface {
	CreateMany(c context.Context, items []ProductAttributeCombination) error
	FetchByID(c context.Context, ID string) (ProductAttributeCombination, error)
	Create(c context.Context, product_attribute_combination *ProductAttributeCombination) error
	Update(c context.Context, product_attribute_combination *ProductAttributeCombination) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeCombination, error)
}
