package domain

import (
	"context"
	directory "earnforglance/server/domain/directory"
	domain "earnforglance/server/domain/localization"
)

type LocalizationRequest struct {
	ID      string
	Filters []Filter
	Sort    string
	Limit   int
	Page    int
	Lang    string
	Rtl     bool
	Content []string
}

type LocalizationResponse struct {
	Language   domain.Language
	Currency   directory.Currency
	Resources  []domain.LocaleStringResource
	Properties []domain.LocalizedProperty
}

type LocalizationsResponse struct {
	Localizations []LocalizationResponse
}

type LocalizationRepository interface {
	GetLocalizations(c context.Context, filter LocalizationRequest) ([]LocalizationsResponse, error)
}

type LocalizationUsecase interface {
	GetLocalizations(c context.Context, filter LocalizationRequest) ([]LocalizationsResponse, error)
}
