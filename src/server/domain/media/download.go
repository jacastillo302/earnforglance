package domain

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDownload = "downloads"
)

// Download represents a downloads
type Download struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	DownloadGuid   uuid.UUID          `bson:"download_guid"`
	UseDownloadUrl bool               `bson:"use_download_url"`
	DownloadUrl    string             `bson:"download_url"`
	DownloadBinary []byte             `bson:"download_binary"`
	ContentType    string             `bson:"content_type"`
	Filename       string             `bson:"filename"`
	Extension      string             `bson:"extension"`
	IsNew          bool               `bson:"is_new"`
}

// DownloadRepository defines the methods that any
// data storage provider needs to implement to get
// and store downloads
type DownloadRepository interface {
	Create(c context.Context, download *Download) error
	Update(c context.Context, download *Download) error
	Delete(c context.Context, download *Download) error
	Fetch(c context.Context) ([]Download, error)
	FetchByID(c context.Context, downloadID string) (Download, error)
}

// DownloadUsecase defines the methods that any
// business logic provider needs to implement to
// manage downloads
type DownloadUsecase interface {
	FetchByID(c context.Context, downloadID string) (Download, error)
	Create(c context.Context, download *Download) error
	Update(c context.Context, download *Download) error
	Delete(c context.Context, download *Download) error
	Fetch(c context.Context) ([]Download, error)
}
