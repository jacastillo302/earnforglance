package domain

import (
	"context"
)

const (
	CollectionMultiFactorAuthenticationSettings = "multiFactor_authentication_settings"
)

// MultiFactorAuthenticationSettings represents multi-factor authentication settings.
type MultiFactorAuthenticationSettings struct {
	// ActiveAuthenticationMethodSystemNames contains system names of active multi-factor authentication methods.
	ActiveAuthenticationMethodSystemNames []string `bson:"activeAuthenticationMethodSystemNames"`

	// ForceMultifactorAuthentication indicates whether to force multi-factor authentication.
	ForceMultifactorAuthentication bool `bson:"forceMultifactorAuthentication"`
}

// NewMultiFactorAuthenticationSettings initializes a new instance of MultiFactorAuthenticationSettings.
func NewMultiFactorAuthenticationSettings() *MultiFactorAuthenticationSettings {
	return &MultiFactorAuthenticationSettings{
		ActiveAuthenticationMethodSystemNames: []string{},
	}
}

type MultiFactorAuthenticationSettingsRepository interface {
	CreateMany(c context.Context, items []MultiFactorAuthenticationSettings) error
	Create(c context.Context, permission_record_customer_role_mapping *MultiFactorAuthenticationSettings) error
	Update(c context.Context, permission_record_customer_role_mapping *MultiFactorAuthenticationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MultiFactorAuthenticationSettings, error)
	FetchByID(c context.Context, ID string) (MultiFactorAuthenticationSettings, error)
}

type MultiFactorAuthenticationSettingsUsecase interface {
	CreateMany(c context.Context, items []MultiFactorAuthenticationSettings) error
	FetchByID(c context.Context, ID string) (MultiFactorAuthenticationSettings, error)
	Create(c context.Context, permission_record_customer_role_mapping *MultiFactorAuthenticationSettings) error
	Update(c context.Context, permission_record_customer_role_mapping *MultiFactorAuthenticationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MultiFactorAuthenticationSettings, error)
}
