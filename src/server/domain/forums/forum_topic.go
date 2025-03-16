package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumTopic = "forum_topics"
)

// ForumTopic represents a forum topic
type ForumTopic struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ForumID            int                `bson:"forum_id"`
	CustomerID         int                `bson:"customer_id"`
	TopicTypeID        int                `bson:"topic_type_id"`
	Subject            string             `bson:"subject"`
	NumPosts           int                `bson:"num_posts"`
	Views              int                `bson:"views"`
	LastPostID         int                `bson:"last_post_id"`
	LastPostCustomerID int                `bson:"last_post_customer_id"`
	LastPostTime       *time.Time         `bson:"last_post_time,omitempty"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc       time.Time          `bson:"updated_on_utc"`
	ForumTopicType     ForumTopicType     `bson:"forum_topic_type"`
}
