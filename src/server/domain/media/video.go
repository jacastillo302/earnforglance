package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionVideo = "videos"
)

// Video represents a videos
type Video struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	VideoUrl string             `bson:"video_url"`
}

// VideoRepository represents the video repository interface
type VideoRepository interface {
	CreateMany(c context.Context, items []Video) error
	Create(c context.Context, video *Video) error
	Update(c context.Context, video *Video) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Video, error)
	FetchByID(c context.Context, ID string) (Video, error)
}

// VideoUsecase represents the video usecase interface
type VideoUsecase interface {
	CreateMany(c context.Context, items []Video) error
	FetchByID(c context.Context, ID string) (Video, error)
	Create(c context.Context, video *Video) error
	Update(c context.Context, video *Video) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Video, error)
}
