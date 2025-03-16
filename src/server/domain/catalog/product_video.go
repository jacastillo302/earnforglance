package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductVideo = "product_videos"
)

// ProductVideo represents a product video mapping
type ProductVideo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    int                `bson:"product_id"`
	VideoID      int                `bson:"video_id"`
	DisplayOrder int                `bson:"display_order"`
}
