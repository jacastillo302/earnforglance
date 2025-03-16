package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPictureBinary = "picture_binaries"
)

// PictureBinary represents a picture binary data
type PictureBinary struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BinaryData []byte             `bson:"binary_data"`
	PictureID  int                `bson:"picture_id"`
}
