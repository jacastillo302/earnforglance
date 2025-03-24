package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductVideo = "product_videos"
)

// ProductVideo represents a product video mapping
type ProductVideo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    primitive.ObjectID `bson:"product_id"`
	VideoID      primitive.ObjectID `bson:"video_id"`
	DisplayOrder int                `bson:"display_order"`
}

type ProductVideoRepository interface {
	CreateMany(c context.Context, items []ProductVideo) error
	Create(c context.Context, product_video *ProductVideo) error
	Update(c context.Context, product_video *ProductVideo) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductVideo, error)
	FetchByID(c context.Context, ID string) (ProductVideo, error)
}

type ProductVideoUsecase interface {
	CreateMany(c context.Context, items []ProductVideo) error
	FetchByID(c context.Context, ID string) (ProductVideo, error)
	Create(c context.Context, product_video *ProductVideo) error
	Update(c context.Context, product_video *ProductVideo) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductVideo, error)
}
