package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumTopic = "forum_topics"
)

// ForumTopic represents a forum topic
type ForumTopic struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ForumID            primitive.ObjectID `bson:"forum_id"`
	CustomerID         primitive.ObjectID `bson:"customer_id"`
	TopicTypeID        int                `bson:"topic_type_id"`
	Subject            string             `bson:"subject"`
	NumPosts           int                `bson:"num_posts"`
	Views              int                `bson:"views"`
	LastPostID         primitive.ObjectID `bson:"last_post_id"`
	LastPostCustomerID primitive.ObjectID `bson:"last_post_customer_id"`
	LastPostTime       *time.Time         `bson:"last_post_time,omitempty"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc       time.Time          `bson:"updated_on_utc"`
	ForumTopicType     ForumTopicType     `bson:"forum_topic_type"`
}

// ForumTopicRepository represents the forum topic repository interface
type ForumTopicRepository interface {
	CreateMany(c context.Context, items []ForumTopic) error
	Create(c context.Context, forum_topic *ForumTopic) error
	Update(c context.Context, forum_topic *ForumTopic) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumTopic, error)
	FetchByID(c context.Context, ID string) (ForumTopic, error)
}

// ForumTopicUsecase represents the forum topic usecase interface
type ForumTopicUsecase interface {
	FetchByID(c context.Context, ID string) (ForumTopic, error)
	Create(c context.Context, forum_topic *ForumTopic) error
	Update(c context.Context, forum_topic *ForumTopic) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumTopic, error)
}
