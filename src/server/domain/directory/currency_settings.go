package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCurrencySettings = "currency_settings"
)

// CurrencySettings represents currency settings
type CurrencySettings struct {
	ID                                   primitive.ObjectID `bson:"_id,omitempty"`
	DisplayCurrencyLabel                 bool               `bson:"display_currency_label"`
	PrimaryStoreCurrencyID               int                `bson:"primary_store_currency_id"`
	PrimaryExchangeRateCurrencyID        int                `bson:"primary_exchange_rate_currency_id"`
	ActiveExchangeRateProviderSystemName string             `bson:"active_exchange_rate_provider_system_name"`
	AutoUpdateEnabled                    bool               `bson:"auto_update_enabled"`
}
