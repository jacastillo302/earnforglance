package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPictureHashes = "picture_hashes"
)

// PictureHashes is a helper class for making picture hashes from DB
type PictureHashes struct {
	PictureID primitive.ObjectID `bson:"picture_id"`
	Hash      []byte             `bson:"hash"`
}

// PictureHashesRepository defines the repository interface for PictureHashes
type PictureHashesRepository interface {
	CreateMany(c context.Context, items []PictureHashes) error
	Create(c context.Context, picture_hashes *PictureHashes) error
	Update(c context.Context, picture_hashes *PictureHashes) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PictureHashes, error)
	FetchByID(c context.Context, ID string) (PictureHashes, error)
}

// PictureHashesUsecase defines the usecase interface for PictureHashes
type PictureHashesUsecase interface {
	FetchByID(c context.Context, ID string) (PictureHashes, error)
	Create(c context.Context, picture_hashes *PictureHashes) error
	Update(c context.Context, picture_hashes *PictureHashes) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PictureHashes, error)
}
