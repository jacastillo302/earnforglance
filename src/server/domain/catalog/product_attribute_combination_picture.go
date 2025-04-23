package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductAttributeCombinationPicture = "product_attribute_combination_pictures"
)

// ProductAttributeCombinationPicture represents a product attribute combination picture
type ProductAttributeCombinationPicture struct {
	ID                            bson.ObjectID `bson:"_id,omitempty"`
	ProductAttributeCombinationID bson.ObjectID `bson:"product_attribute_combination_id"`
	PictureID                     bson.ObjectID `bson:"picture_id"`
}

type ProductAttributeCombinationPictureRepository interface {
	CreateMany(c context.Context, items []ProductAttributeCombinationPicture) error
	Create(c context.Context, product_attribute_combination_picture *ProductAttributeCombinationPicture) error
	Update(c context.Context, product_attribute_combination_picture *ProductAttributeCombinationPicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeCombinationPicture, error)
	FetchByID(c context.Context, ID string) (ProductAttributeCombinationPicture, error)
}

type ProductAttributeCombinationPictureUsecase interface {
	CreateMany(c context.Context, items []ProductAttributeCombinationPicture) error
	FetchByID(c context.Context, ID string) (ProductAttributeCombinationPicture, error)
	Create(c context.Context, product_attribute_combination_picture *ProductAttributeCombinationPicture) error
	Update(c context.Context, product_attribute_combination_picture *ProductAttributeCombinationPicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductAttributeCombinationPicture, error)
}
