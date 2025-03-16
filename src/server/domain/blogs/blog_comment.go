package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlogComment = "blog_comments"
)

// BlogComment represents a blog comment
type BlogComment struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID   int                `bson:"customer_id"`
	CommentText  string             `bson:"comment_text"`
	IsApproved   bool               `bson:"is_approved"`
	StoreID      int                `bson:"store_id"`
	BlogPostID   int                `bson:"blog_post_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
