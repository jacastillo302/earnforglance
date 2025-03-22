package domain

import (
	"context" // added context library
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
	NewsItemID   primitive.ObjectID `bson:"news_item_id"`
	CustomerID   primitive.ObjectID `bson:"customer_id"`
	IsApproved   bool               `bson:"is_approved"`
	StoreID      primitive.ObjectID `bson:"store_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}

// NewsCommentRepository interface
type NewsCommentRepository interface {
	Create(c context.Context, news_comment *NewsComment) error
	Update(c context.Context, news_comment *NewsComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsComment, error)
	FetchByID(c context.Context, ID string) (NewsComment, error)
}

// NewsCommentUsecase interface
type NewsCommentUsecase interface {
	FetchByID(c context.Context, ID string) (NewsComment, error)
	Create(c context.Context, news_comment *NewsComment) error
	Update(c context.Context, news_comment *NewsComment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsComment, error)
}
