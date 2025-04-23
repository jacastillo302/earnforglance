package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionNewsComment = "news_comments"
)

// NewsComment represents a news comment
type NewsComment struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	CommentTitle string        `bson:"comment_title"`
	CommentText  string        `bson:"comment_text"`
	NewsItemID   bson.ObjectID `bson:"news_item_id"`
	CustomerID   bson.ObjectID `bson:"customer_id"`
	IsApproved   bool          `bson:"is_approved"`
	StoreID      bson.ObjectID `bson:"store_id"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

// NewsCommentRepository interface
type NewsCommentRepository interface {
	CreateMany(c context.Context, items []NewsComment) error
	Create(c context.Context, news_comment *NewsComment) error
	Update(c context.Context, news_comment *NewsComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsComment, error)
	FetchByID(c context.Context, ID string) (NewsComment, error)
}

// NewsCommentUsecase interface
type NewsCommentUsecase interface {
	CreateMany(c context.Context, items []NewsComment) error
	FetchByID(c context.Context, ID string) (NewsComment, error)
	Create(c context.Context, news_comment *NewsComment) error
	Update(c context.Context, news_comment *NewsComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsComment, error)
}
