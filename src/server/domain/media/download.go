package domain

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDownload = "downloads"
)

// Download represents a download
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
