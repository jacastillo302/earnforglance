package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPictureBinary = "picture_binaries"
)

// PictureBinary represents a picture binary data all
type PictureBinary struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BinaryData []byte             `bson:"binary_data"`
	PictureID  primitive.ObjectID `bson:"picture_id"`
}

// PictureBinaryRepository defines the repository interface for PictureBinary
type PictureBinaryRepository interface {
	CreateMany(c context.Context, items []PictureBinary) error
	Create(c context.Context, picture_binary *PictureBinary) error
	Update(c context.Context, picture_binary *PictureBinary) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PictureBinary, error)
	FetchByID(c context.Context, ID string) (PictureBinary, error)
}

// PictureBinaryUsecase defines the usecase interface for PictureBinary
type PictureBinaryUsecase interface {
	CreateMany(c context.Context, items []PictureBinary) error
	FetchByID(c context.Context, ID string) (PictureBinary, error)
	Create(c context.Context, picture_binary *PictureBinary) error
	Update(c context.Context, picture_binary *PictureBinary) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PictureBinary, error)
}
