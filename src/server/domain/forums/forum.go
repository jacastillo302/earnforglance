package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForum = "forums"
)

// Forum represents a forum
type Forum struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ForumGroupID       int                `bson:"forum_group_id"`
	Name               string             `bson:"name"`
	Description        string             `bson:"description"`
	NumTopics          int                `bson:"num_topics"`
	NumPosts           int                `bson:"num_posts"`
	LastTopicID        int                `bson:"last_topic_id"`
	LastPostID         int                `bson:"last_post_id"`
	LastPostCustomerID int                `bson:"last_post_customer_id"`
	LastPostTime       *time.Time         `bson:"last_post_time,omitempty"`
	DisplayOrder       int                `bson:"display_order"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc       time.Time          `bson:"updated_on_utc"`
}
