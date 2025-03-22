package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCrossSellProduct = "cross_sell_products"
)

// CrossSellProduct represents a cross-sell product
type CrossSellProduct struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProductID1 primitive.ObjectID `bson:"product_id1"`
	ProductID2 primitive.ObjectID `bson:"product_id2"`
}

type CrossSellProductRepository interface {
	Create(c context.Context, cross_sell_product *CrossSellProduct) error
	Update(c context.Context, cross_sell_product *CrossSellProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CrossSellProduct, error)
	FetchByID(c context.Context, ID string) (CrossSellProduct, error)
}

type CrossSellProductUsecase interface {
	FetchByID(c context.Context, ID string) (CrossSellProduct, error)
	Create(c context.Context, cross_sell_product *CrossSellProduct) error
	Update(c context.Context, cross_sell_product *CrossSellProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CrossSellProduct, error)
}
