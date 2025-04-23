package domain

import (
	"context" // added context library
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionNewsLetterSubscription = "news_letter_subscriptions"
)

// NewsLetterSubscription represents NewsLetterSubscription entity
type NewsLetterSubscription struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Guid         uuid.UUID     `bson:"newsletter_subscription_guid"`
	Email        string        `bson:"email"`
	Active       bool          `bson:"active"`
	StoreID      bson.ObjectID `bson:"store_id"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
	IpAddress    string        `bson:"last_ip_address"`
	LanguageID   bson.ObjectID `bson:"language_id"`
}

// NewsLetterSubscriptionRepository interface
type NewsLetterSubscriptionRepository interface {
	CreateMany(c context.Context, items []NewsLetterSubscription) error
	Create(c context.Context, newsletter_subscription *NewsLetterSubscription) error
	Update(c context.Context, newsletter_subscription *NewsLetterSubscription) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsLetterSubscription, error)
	FetchByID(c context.Context, ID string) (NewsLetterSubscription, error)
}

// NewsLetterSubscriptionUsecase interface
type NewsLetterSubscriptionUsecase interface {
	CreateMany(c context.Context, items []NewsLetterSubscription) error
	FetchByID(c context.Context, ID string) (NewsLetterSubscription, error)
	Create(c context.Context, newsletter_subscription *NewsLetterSubscription) error
	Update(c context.Context, newsletter_subscription *NewsLetterSubscription) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]NewsLetterSubscription, error)
}
