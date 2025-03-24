package domain

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionExchangeRate = "exchange_rates"
)

// ExchangeRate represents an exchange rate
type ExchangeRate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CurrencyCode string             `bson:"currency_code"`
	Rate         float64            `bson:"rate"`
	UpdatedOn    time.Time          `bson:"updated_on"`
}

type ExchangeRateRepository interface {
	CreateMany(c context.Context, items []ExchangeRate) error
	Create(c context.Context, exchange_rate *ExchangeRate) error
	Update(c context.Context, exchange_rate *ExchangeRate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ExchangeRate, error)
	FetchByID(c context.Context, ID string) (ExchangeRate, error)
}

type ExchangeRateUsecase interface {
	FetchByID(c context.Context, ID string) (ExchangeRate, error)
	Create(c context.Context, exchange_rate *ExchangeRate) error
	Update(c context.Context, exchange_rate *ExchangeRate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ExchangeRate, error)
}

// NewExchangeRate creates a new instance of the ExchangeRate struct
func NewExchangeRate() *ExchangeRate {
	return &ExchangeRate{
		CurrencyCode: "",
		Rate:         1.0,
	}
}

// ToString formats the rate into a string with the currency code, e.g. "USD 0.72543"
func (e *ExchangeRate) ToString() string {
	return fmt.Sprintf("%s %.5f", e.CurrencyCode, e.Rate)
}
