package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionForumTopic = "forum_topics"
)

// ForumTopic represents a forum topic
type ForumTopic struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	ForumID            bson.ObjectID `bson:"forum_id"`
	CustomerID         bson.ObjectID `bson:"customer_id"`
	TopicTypeID        int           `bson:"topic_type_id"`
	Subject            string        `bson:"subject"`
	NumPosts           int           `bson:"num_posts"`
	Views              int           `bson:"views"`
	LastPostID         bson.ObjectID `bson:"last_post_id"`
	LastPostCustomerID bson.ObjectID `bson:"last_post_customer_id"`
	LastPostTime       *time.Time    `bson:"last_post_time"`
	CreatedOnUtc       time.Time     `bson:"created_on_utc"`
	UpdatedOnUtc       time.Time     `bson:"updated_on_utc"`
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
	CreateMany(c context.Context, items []ForumTopic) error
	FetchByID(c context.Context, ID string) (ForumTopic, error)
	Create(c context.Context, forum_topic *ForumTopic) error
	Update(c context.Context, forum_topic *ForumTopic) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ForumTopic, error)
}
