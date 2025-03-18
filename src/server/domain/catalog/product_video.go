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
	Create(c context.Context, product_video *ProductVideo) error
	Update(c context.Context, product_video *ProductVideo) error
	Delete(c context.Context, product_video *ProductVideo) error
	Fetch(c context.Context) ([]ProductVideo, error)
	FetchByID(c context.Context, product_videoID string) (ProductVideo, error)
}

type ProductVideoUsecase interface {
	FetchByID(c context.Context, product_videoID string) (ProductVideo, error)
	Create(c context.Context, product_video *ProductVideo) error
	Update(c context.Context, product_video *ProductVideo) error
	Delete(c context.Context, product_video *ProductVideo) error
	Fetch(c context.Context) ([]ProductVideo, error)
}
