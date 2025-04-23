package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductAttributeValuePicture = "product_attribute_value_pictures"
)

// ProductAttributeValuePicture represents a product attribute value picture
type ProductAttributeValuePicture struct {
	ID                      bson.ObjectID `bson:"_id,omitempty"`
	ProductAttributeValueID bson.ObjectID `bson:"product_attribute_value_id"`
	PictureID               bson.ObjectID `bson:"picture_id"`
}

type ProductAttributeValuePictureRepository interface {
	CreateMany(c context.Context, items []ProductAttributeValuePicture) error
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValuePicture, error)
	FetchByID(c context.Context, ID string) (ProductAttributeValuePicture, error)
}

type ProductAttributeValuePictureUsecase interface {
	CreateMany(c context.Context, items []ProductAttributeValuePicture) error
	FetchByID(c context.Context, ID string) (ProductAttributeValuePicture, error)
	Create(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Update(c context.Context, product_attribute_value_picture *ProductAttributeValuePicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeValuePicture, error)
}
