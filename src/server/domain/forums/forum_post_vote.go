package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumPostVote = "forum_post_votes"
)

// ForumPostVote represents a forum post vote
type ForumPostVote struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ForumPostID  int                `bson:"forum_post_id"`
	CustomerID   int                `bson:"customer_id"`
	IsUp         bool               `bson:"is_up"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
