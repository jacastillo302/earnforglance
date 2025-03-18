package domain

import (
	"context" // added context library
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionForumSubscription = "forum_subscriptions"
)

// ForumSubscription represents a forum subscription item
type ForumSubscription struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SubscriptionGuid uuid.UUID          `bson:"subscription_guid"`
	CustomerID       int                `bson:"customer_id"`
	ForumID          int                `bson:"forum_id"`
	TopicID          int                `bson:"topic_id"`
	CreatedOnUtc     time.Time          `bson:"created_on_utc"`
}

// ForumSubscriptionRepository interface
type ForumSubscriptionRepository interface {
	Create(c context.Context, forum_subscription *ForumSubscription) error
	Update(c context.Context, forum_subscription *ForumSubscription) error
	Delete(c context.Context, forum_subscription *ForumSubscription) error
	Fetch(c context.Context) ([]ForumSubscription, error)
	FetchByID(c context.Context, forum_subscriptionID string) (ForumSubscription, error)
}

// ForumSubscriptionUsecase interface
type ForumSubscriptionUsecase interface {
	FetchByID(c context.Context, forum_subscriptionID string) (ForumSubscription, error)
	Create(c context.Context, forum_subscription *ForumSubscription) error
	Update(c context.Context, forum_subscription *ForumSubscription) error
	Delete(c context.Context, forum_subscription *ForumSubscription) error
	Fetch(c context.Context) ([]ForumSubscription, error)
}
