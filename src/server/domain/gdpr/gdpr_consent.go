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
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`
	Title                     string              `bson:"title"`
	Message                   string              `bson:"message"`
	ParentID                  *primitive.ObjectID `bson:"parent_id"`
	IsRequired                bool                `bson:"is_required"`
	RequiredMessage           string              `bson:"required_message"`
	DisplayDuringRegistration bool                `bson:"display_during_registration"`
	DisplayOnCustomerInfoPage bool                `bson:"display_on_customer_info_page"`
	DisplayOrder              int                 `bson:"display_order"`
}

// GdprConsentRepository interface
type GdprConsentRepository interface {
	CreateMany(c context.Context, items []GdprConsent) error
	Create(c context.Context, gdpr_consent *GdprConsent) error
	Update(c context.Context, gdpr_consent *GdprConsent) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprConsent, error)
	FetchByID(c context.Context, ID string) (GdprConsent, error)
}

// GdprConsentUsecase interface
type GdprConsentUsecase interface {
	CreateMany(c context.Context, items []GdprConsent) error
	FetchByID(c context.Context, ID string) (GdprConsent, error)
	Create(c context.Context, gdpr_consent *GdprConsent) error
	Update(c context.Context, gdpr_consent *GdprConsent) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprConsent, error)
}
