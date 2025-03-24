package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCurrencySettings = "currency_settings"
)

// CurrencySettings represents currency settings
type CurrencySettings struct {
	ID                                   primitive.ObjectID `bson:"_id,omitempty"`
	DisplayCurrencyLabel                 bool               `bson:"display_currency_label"`
	PrimaryStoreCurrencyID               primitive.ObjectID `bson:"primary_store_currency_id"`
	PrimaryExchangeRateCurrencyID        primitive.ObjectID `bson:"primary_exchange_rate_currency_id"`
	ActiveExchangeRateProviderSystemName string             `bson:"active_exchange_rate_provider_system_name"`
	AutoUpdateEnabled                    bool               `bson:"auto_update_enabled"`
}

type CurrencySettingsRepository interface {
	CreateMany(c context.Context, items []CurrencySettings) error
	Create(c context.Context, currency_settings *CurrencySettings) error
	Update(c context.Context, currency_settings *CurrencySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CurrencySettings, error)
	FetchByID(c context.Context, ID string) (CurrencySettings, error)
}

type CurrencySettingsUsecase interface {
	FetchByID(c context.Context, ID string) (CurrencySettings, error)
	Create(c context.Context, currency_settings *CurrencySettings) error
	Update(c context.Context, currency_settings *CurrencySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CurrencySettings, error)
}
