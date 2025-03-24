package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumPost = "forum_posts"
)

// ForumPost represents a forum post
type ForumPost struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	TopicID      primitive.ObjectID `bson:"topic_id"`
	CustomerID   primitive.ObjectID `bson:"customer_id"`
	Text         string             `bson:"text"`
	IPAddress    string             `bson:"ip_address"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc time.Time          `bson:"updated_on_utc"`
	VoteCount    int                `bson:"vote_count"`
}

// ForumPostRepository interface
type ForumPostRepository interface {
	CreateMany(c context.Context, items []ForumPost) error
	Create(c context.Context, forum_post *ForumPost) error
	Update(c context.Context, forum_post *ForumPost) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumPost, error)
	FetchByID(c context.Context, ID string) (ForumPost, error)
}

// ForumPostUsecase interface
type ForumPostUsecase interface {
	FetchByID(c context.Context, ID string) (ForumPost, error)
	Create(c context.Context, forum_post *ForumPost) error
	Update(c context.Context, forum_post *ForumPost) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumPost, error)
}
