package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductAttributeCombinationPicture = "product_attribute_combination_pictures"
)

// ProductAttributeCombinationPicture represents a product attribute combination picture
type ProductAttributeCombinationPicture struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`
	ProductAttributeCombinationID int                `bson:"product_attribute_combination_id"`
	PictureID                     int                `bson:"picture_id"`
}
