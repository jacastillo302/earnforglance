package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductProductTagMapping = "product_product_tag_mappings"
)

// ProductProductTagMapping represents a product-product tag mapping class
type ProductProductTagMapping struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    int                `bson:"product_id"`
	ProductTagID int                `bson:"product_tag_id"`
}
