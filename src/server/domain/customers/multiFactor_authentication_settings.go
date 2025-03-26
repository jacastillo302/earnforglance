package domain

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
