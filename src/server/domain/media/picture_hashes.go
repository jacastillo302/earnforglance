package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPictureHash = "picture_hashes"
)

// PictureHashItem is a helper class for making picture hashes from DB
type PictureHashItem struct {
	PictureID primitive.ObjectID `bson:"picture_id"`
	Hash      []byte             `bson:"hash"`
}

// PictureHashItemRepository defines the repository interface for PictureHashItem
type PictureHashItemRepository interface {
	Create(c context.Context, picture_hashes *PictureHashItem) error
	Update(c context.Context, picture_hashes *PictureHashItem) error
	Delete(c context.Context, picture_hashes *PictureHashItem) error
	Fetch(c context.Context) ([]PictureHashItem, error)
	FetchByID(c context.Context, picture_hashesID string) (PictureHashItem, error)
}

// PictureHashItemUsecase defines the usecase interface for PictureHashItem
type PictureHashItemUsecase interface {
	FetchByID(c context.Context, picture_hashesID string) (PictureHashItem, error)
	Create(c context.Context, picture_hashes *PictureHashItem) error
	Update(c context.Context, picture_hashes *PictureHashItem) error
	Delete(c context.Context, picture_hashes *PictureHashItem) error
	Fetch(c context.Context) ([]PictureHashItem, error)
}
