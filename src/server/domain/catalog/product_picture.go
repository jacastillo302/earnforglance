package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductPicture = "product_pictures"
)

// ProductPicture represents a product picture mapping
type ProductPicture struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    int                `bson:"product_id"`
	PictureID    int                `bson:"picture_id"`
	DisplayOrder int                `bson:"display_order"`
}
