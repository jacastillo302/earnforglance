package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGdprSettings = "gdpr_settings"
)

// GdprSettings represents GDPR settings all
type GdprSettings struct {
	ID                                 primitive.ObjectID `bson:"_id,omitempty"`
	GdprEnabled                        bool               `bson:"gdpr_enabled"`
	LogPrivacyPolicyConsent            bool               `bson:"log_privacy_policy_consent"`
	LogNewsletterConsent               bool               `bson:"log_newsletter_consent"`
	LogUserProfileChanges              bool               `bson:"log_user_profile_changes"`
	DeleteInactiveCustomersAfterMonths int                `bson:"delete_inactive_customers_after_months"`
}

type GdprSettingsRepository interface {
	CreateMany(c context.Context, items []GdprSettings) error
	Create(c context.Context, gdpr_settings *GdprSettings) error
	Update(c context.Context, gdpr_settings *GdprSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprSettings, error)
	FetchByID(c context.Context, ID string) (GdprSettings, error)
}

type GdprSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (GdprSettings, error)
	Create(c context.Context, gdpr_settings *GdprSettings) error
	Update(c context.Context, gdpr_settings *GdprSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprSettings, error)
}
