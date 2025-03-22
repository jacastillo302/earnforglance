package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForum = "forums"
)

// Forum represents a forum
type Forum struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ForumGroupID       primitive.ObjectID `bson:"forum_group_id"`
	Name               string             `bson:"name"`
	Description        string             `bson:"description"`
	NumTopics          int                `bson:"num_topics"`
	NumPosts           int                `bson:"num_posts"`
	LastTopicID        primitive.ObjectID `bson:"last_topic_id"`
	LastPostID         primitive.ObjectID `bson:"last_post_id"`
	LastPostCustomerID primitive.ObjectID `bson:"last_post_customer_id"`
	LastPostTime       *time.Time         `bson:"last_post_time,omitempty"`
	DisplayOrder       int                `bson:"display_order"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc       time.Time          `bson:"updated_on_utc"`
}

// ForumRepository interface
type ForumRepository interface {
	Create(c context.Context, forum *Forum) error
	Update(c context.Context, forum *Forum) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Forum, error)
	FetchByID(c context.Context, ID string) (Forum, error)
}

// ForumUsecase interface
type ForumUsecase interface {
	FetchByID(c context.Context, ID string) (Forum, error)
	Create(c context.Context, forum *Forum) error
	Update(c context.Context, forum *Forum) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Forum, error)
}
