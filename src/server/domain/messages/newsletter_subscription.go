package domain

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNewsLetterSubscriptiont = "newsletter_subscriptions"
)

// NewsLetterSubscription represents NewsLetterSubscription entity
type NewsLetterSubscription struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	NewsLetterSubscriptionGuid uuid.UUID          `bson:"newsletter_subscription_guid"`
	Email                      string             `bson:"email"`
	Active                     bool               `bson:"active"`
	StoreID                    int                `bson:"store_id"`
	CreatedOnUtc               time.Time          `bson:"created_on_utc"`
	LanguageID                 int                `bson:"language_id"`
}
