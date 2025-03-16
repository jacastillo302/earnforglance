package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCrossSellProduct = "cross_sell_products"
)

// CrossSellProduct represents a cross-sell product
type CrossSellProduct struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProductID1 int                `bson:"product_id1"`
	ProductID2 int                `bson:"product_id2"`
}
