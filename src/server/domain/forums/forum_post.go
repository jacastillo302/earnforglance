package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumPost = "forum_posts"
)

// ForumPost represents a forum post
type ForumPost struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	TopicID      int                `bson:"topic_id"`
	CustomerID   int                `bson:"customer_id"`
	Text         string             `bson:"text"`
	IPAddress    string             `bson:"ip_address"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc time.Time          `bson:"updated_on_utc"`
	VoteCount    int                `bson:"vote_count"`
}
