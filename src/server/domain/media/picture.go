package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPicture = "pictures"
)

// Picture represents a picture
type Picture struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	MimeType       string             `bson:"mime_type"`
	SeoFilename    string             `bson:"seo_filename"`
	AltAttribute   string             `bson:"alt_attribute"`
	TitleAttribute string             `bson:"title_attribute"`
	IsNew          bool               `bson:"is_new"`
	VirtualPath    string             `bson:"virtual_path"`
}
