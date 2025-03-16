package domain

import (
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
