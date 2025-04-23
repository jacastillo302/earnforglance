package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCrossSellProduct = "cross_sell_products"
)

// CrossSellProduct represents a cross-sell product
type CrossSellProduct struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	ProductID1 bson.ObjectID `bson:"product_id1"`
	ProductID2 bson.ObjectID `bson:"product_id2"`
}

type CrossSellProductRepository interface {
	CreateMany(c context.Context, items []CrossSellProduct) error
	Create(c context.Context, cross_sell_product *CrossSellProduct) error
	Update(c context.Context, cross_sell_product *CrossSellProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CrossSellProduct, error)
	FetchByID(c context.Context, ID string) (CrossSellProduct, error)
}

type CrossSellProductUsecase interface {
	CreateMany(c context.Context, items []CrossSellProduct) error
	FetchByID(c context.Context, ID string) (CrossSellProduct, error)
	Create(c context.Context, cross_sell_product *CrossSellProduct) error
	Update(c context.Context, cross_sell_product *CrossSellProduct) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CrossSellProduct, error)
}
