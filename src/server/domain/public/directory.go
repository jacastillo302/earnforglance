package domain

import (
	"context"
	domain "earnforglance/server/domain/directory"
)

type CountryRequest struct {
	ID             string
	Filters        []Filter
	Sort           string
	AllowsBilling  bool
	AllowsShipping bool
	SubjectToVat   bool
	Content        []string
}

type CountryResponse struct {
	Countries domain.Country
	States    []domain.StateProvince
}

type CountriesResponse struct {
	Countries []CountryResponse
}

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
	GetCountries(c context.Context, filter CountryRequest) ([]CountriesResponse, error)
}

type DirectoryUsecase interface {
	GetCurrencies(c context.Context, filter CurrencyRequest) ([]CurrenciesResponse, error)
	GetCountries(c context.Context, filter CountryRequest) ([]CountriesResponse, error)
}
