package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGdprConsent = "gdpr_consents"
)

// GdprConsent represents a GDPR consents
type GdprConsent struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	Message                   string             `bson:"message"`
	IsRequired                bool               `bson:"is_required"`
	RequiredMessage           string             `bson:"required_message"`
	DisplayDuringRegistration bool               `bson:"display_during_registration"`
	DisplayOnCustomerInfoPage bool               `bson:"display_on_customer_info_page"`
	DisplayOrder              int                `bson:"display_order"`
}

// GdprConsentRepository interface
type GdprConsentRepository interface {
	Create(c context.Context, gdpr_consent *GdprConsent) error
	Update(c context.Context, gdpr_consent *GdprConsent) error
	Delete(c context.Context, gdpr_consent *GdprConsent) error
	Fetch(c context.Context) ([]GdprConsent, error)
	FetchByID(c context.Context, gdpr_consentID string) (GdprConsent, error)
}

// GdprConsentUsecase interface
type GdprConsentUsecase interface {
	FetchByID(c context.Context, gdpr_consentID string) (GdprConsent, error)
	Create(c context.Context, gdpr_consent *GdprConsent) error
	Update(c context.Context, gdpr_consent *GdprConsent) error
	Delete(c context.Context, gdpr_consent *GdprConsent) error
	Fetch(c context.Context) ([]GdprConsent, error)
}
