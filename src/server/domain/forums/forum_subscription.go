package domain

import (
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
