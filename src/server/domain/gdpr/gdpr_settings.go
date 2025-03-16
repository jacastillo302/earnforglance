package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionGdprSettings = "gdpr_settings"
)

// GdprSettings represents GDPR settings
type GdprSettings struct {
	ID                                 primitive.ObjectID `bson:"_id,omitempty"`
	GdprEnabled                        bool               `bson:"gdpr_enabled"`
	LogPrivacyPolicyConsent            bool               `bson:"log_privacy_policy_consent"`
	LogNewsletterConsent               bool               `bson:"log_newsletter_consent"`
	LogUserProfileChanges              bool               `bson:"log_user_profile_changes"`
	DeleteInactiveCustomersAfterMonths int                `bson:"delete_inactive_customers_after_months"`
}
