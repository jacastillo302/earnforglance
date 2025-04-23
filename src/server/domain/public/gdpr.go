package domain

import (
	"context"
	domain "earnforglance/server/domain/gdpr"
)

type GdprConsentRequest struct {
	ID                        string
	Filters                   []Filter
	Sort                      string
	Limit                     int
	Page                      int
	Lang                      string
	IsRequired                bool
	DisplayDuringRegistration bool
	DisplayOnCustomerInfoPage bool
	Content                   []string
}

type GdprConsentResponse struct {
	GdprConsent domain.GdprConsent
	Parent      *domain.GdprConsent
}

type GdprConsentsResponse struct {
	GdprConsents []GdprConsentResponse
}

type GdprConsentRepository interface {
	GetGdprConsents(c context.Context, filter GdprConsentRequest) ([]GdprConsentsResponse, error)
}

type GdprConsentUsecase interface {
	GetGdprConsents(c context.Context, filter GdprConsentRequest) ([]GdprConsentsResponse, error)
}
