package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPicture = "pictures"
)

// Picture represents a pictures
type Picture struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	MimeType       string        `bson:"mime_type"`
	SeoFilename    string        `bson:"seo_filename"`
	AltAttribute   string        `bson:"alt_attribute"`
	TitleAttribute string        `bson:"title_attribute"`
	IsNew          bool          `bson:"is_new"`
	VirtualPath    string        `bson:"virtual_path"`
}

// PictureRepository represents the picture repository interface
type PictureRepository interface {
	CreateMany(c context.Context, items []Picture) error
	Create(c context.Context, picture *Picture) error
	Update(c context.Context, picture *Picture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Picture, error)
	FetchByID(c context.Context, ID string) (Picture, error)
}

// PictureUsecase represents the picture usecase interface
type PictureUsecase interface {
	CreateMany(c context.Context, items []Picture) error
	FetchByID(c context.Context, ID string) (Picture, error)
	Create(c context.Context, picture *Picture) error
	Update(c context.Context, picture *Picture) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Picture, error)
}
