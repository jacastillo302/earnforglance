package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNewsComment = "news_comments"
)

// NewsComment represents a news comment
type NewsComment struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CommentTitle string             `bson:"comment_title"`
	CommentText  string             `bson:"comment_text"`
	NewsItemID   int                `bson:"news_item_id"`
	CustomerID   int                `bson:"customer_id"`
	IsApproved   bool               `bson:"is_approved"`
	StoreID      int                `bson:"store_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
