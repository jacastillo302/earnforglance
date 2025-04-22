package domain

import (
	"context"
	domain "earnforglance/server/domain/directory"
)

type CurrencyRequest struct {
	ID      string
	Filters []Filter
	Sort    string
	Content []string
}

type CurrencyResponse struct {
	Currency     domain.Currency
	ExchangeRate []domain.ExchangeRate
	RoundingType RoundingType
}

type CurrenciesResponse struct {
	Currencies []CurrencyResponse
}

type RoundingType struct {
	Name        string
	Value       int
	Description string
}

type DirectoryRepository interface {
	GetCurrencies(c context.Context, filter CurrencyRequest) ([]CurrenciesResponse, error)
}

type DirectoryUsecase interface {
	GetCurrencies(c context.Context, filter CurrencyRequest) ([]CurrenciesResponse, error)
}
