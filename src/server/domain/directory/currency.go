package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCurrency = "currencies"
)

// Currency represents a currency
type Currency struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Name             string        `bson:"name"`
	CurrencyCode     string        `bson:"currency_code"`
	Rate             float64       `bson:"rate"`
	DisplayLocale    string        `bson:"display_locale"`
	CustomFormatting string        `bson:"custom_formatting"`
	LimitedToStores  bool          `bson:"limited_to_stores"`
	Published        bool          `bson:"published"`
	DisplayOrder     int           `bson:"display_order"`
	CreatedOnUtc     time.Time     `bson:"created_on_utc"`
	UpdatedOnUtc     time.Time     `bson:"updated_on_utc"`
	RoundingTypeID   int           `bson:"rounding_type_id"`
}

type CurrencyRepository interface {
	CreateMany(c context.Context, items []Currency) error
	Create(c context.Context, currency *Currency) error
	Update(c context.Context, currency *Currency) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Currency, error)
	FetchByID(c context.Context, ID string) (Currency, error)
}

type CurrencyUsecase interface {
	CreateMany(c context.Context, items []Currency) error
	FetchByID(c context.Context, ID string) (Currency, error)
	Create(c context.Context, currency *Currency) error
	Update(c context.Context, currency *Currency) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Currency, error)
}
