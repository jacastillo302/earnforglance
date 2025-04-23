package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductPicture = "product_pictures"
)

// ProductPicture represents a product picture mapping
type ProductPicture struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	ProductID    bson.ObjectID `bson:"product_id"`
	PictureID    bson.ObjectID `bson:"picture_id"`
	DisplayOrder int           `bson:"display_order"`
}

type ProductPictureRepository interface {
	CreateMany(c context.Context, items []ProductPicture) error
	Create(c context.Context, product_picture *ProductPicture) error
	Update(c context.Context, product_picture *ProductPicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductPicture, error)
	FetchByID(c context.Context, ID string) (ProductPicture, error)
}

type ProductPictureUsecase interface {
	CreateMany(c context.Context, items []ProductPicture) error
	FetchByID(c context.Context, ID string) (ProductPicture, error)
	Create(c context.Context, product_picture *ProductPicture) error
	Update(c context.Context, product_picture *ProductPicture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductPicture, error)
}
