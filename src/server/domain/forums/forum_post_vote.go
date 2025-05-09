package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionForumPostVote = "forum_post_votes"
)

// ForumPostVote represents a forum post vote
type ForumPostVote struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	ForumPostID  bson.ObjectID `bson:"forum_post_id"`
	CustomerID   bson.ObjectID `bson:"customer_id"`
	IsUp         bool          `bson:"is_up"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

type ForumPostVoteRepository interface {
	CreateMany(c context.Context, items []ForumPostVote) error
	Create(c context.Context, product_tag *ForumPostVote) error
	Update(c context.Context, product_tag *ForumPostVote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumPostVote, error)
	FetchByID(c context.Context, ID string) (ForumPostVote, error)
}

type ForumPostVoteUsecase interface {
	CreateMany(c context.Context, items []ForumPostVote) error
	FetchByID(c context.Context, ID string) (ForumPostVote, error)
	Create(c context.Context, product_tag *ForumPostVote) error
	Update(c context.Context, product_tag *ForumPostVote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumPostVote, error)
}
