package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductAttributeValuePicture = "product_attribute_value_pictures"
)

// ProductAttributeValuePicture represents a product attribute value picture
type ProductAttributeValuePicture struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	ProductAttributeValueID int                `bson:"product_attribute_value_id"`
	PictureID               int                `bson:"picture_id"`
}
