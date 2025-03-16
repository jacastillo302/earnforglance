package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionVideo = "videos"
)

// Video represents a video
type Video struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	VideoUrl string             `bson:"video_url"`
}
