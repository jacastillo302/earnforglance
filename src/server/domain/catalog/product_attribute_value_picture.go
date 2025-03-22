package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductAttributeValuePicture = "product_attribute_value_pictures"
)

// ProductAttributeValuePicture represents a product attribute value picture
type ProductAttributeValuePicture struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	ProductAttributeValueID primitive.ObjectID `bson:"product_attribute_value_id"`
	PictureID               primitive.ObjectID `bson:"picture_id"`
}

type ProductAttributeValuePictureRepository interface {
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValuePicture, error)
	FetchByID(c context.Context, ID string) (ProductAttributeValuePicture, error)
}

type ProductAttributeValuePictureUsecase interface {
	FetchByID(c context.Context, ID string) (ProductAttributeValuePicture, error)
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValuePicture, error)
}
