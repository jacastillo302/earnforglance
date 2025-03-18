package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionEmailAccountSettings = "email_account_settings"
)

// EmailAccountSettings  represents email account settings
type EmailAccountSettings struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	DefaultEmailAccountID primitive.ObjectID `bson:"default_email_account_id"`
}

// EmailAccountSettingsRepository represents the repository interface for EmailAccountSettings
type EmailAccountSettingsRepository interface {
	Create(c context.Context, email_account_settings *EmailAccountSettings) error
	Update(c context.Context, email_account_settings *EmailAccountSettings) error
	Delete(c context.Context, email_account_settings *EmailAccountSettings) error
	Fetch(c context.Context) ([]EmailAccountSettings, error)
	FetchByID(c context.Context, email_account_settingsID string) (EmailAccountSettings, error)
}

// EmailAccountSettingsUsecase represents the use case interface for EmailAccountSettings
type EmailAccountSettingsUsecase interface {
	FetchByID(c context.Context, email_account_settingsID string) (EmailAccountSettings, error)
	Create(c context.Context, email_account_settings *EmailAccountSettings) error
	Update(c context.Context, email_account_settings *EmailAccountSettings) error
	Delete(c context.Context, email_account_settings *EmailAccountSettings) error
	Fetch(c context.Context) ([]EmailAccountSettings, error)
}
