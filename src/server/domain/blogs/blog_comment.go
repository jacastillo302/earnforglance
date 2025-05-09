package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionBlogComment = "blog_comments"
)

// BlogComment represents a blog comment
type BlogComment struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	CustomerID   bson.ObjectID `bson:"customer_id"`
	CommentText  string        `bson:"comment_text"`
	IsApproved   bool          `bson:"is_approved"`
	StoreID      bson.ObjectID `bson:"store_id"`
	BlogPostID   bson.ObjectID `bson:"blog_post_id"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

type BlogCommentRepository interface {
	CreateMany(c context.Context, items []BlogComment) error
	Create(c context.Context, blog_comment *BlogComment) error
	Update(c context.Context, blog_comment *BlogComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogComment, error)
	FetchByID(c context.Context, blog_commentID string) (BlogComment, error)
}

type BlogCommentUsecase interface {
	CreateMany(c context.Context, items []BlogComment) error
	FetchByID(c context.Context, blog_commentID string) (BlogComment, error)
	Create(c context.Context, blog_comment *BlogComment) error
	Update(c context.Context, blog_comment *BlogComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BlogComment, error)
}
