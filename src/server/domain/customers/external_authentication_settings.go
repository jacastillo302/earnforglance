package domain

const (
	CollectionExternalAuthenticationSettings = "external_authentication_settings"
)

// ExternalAuthenticationSettings represents external authentication settings.
type ExternalAuthenticationSettings struct {
	// RequireEmailValidation indicates whether email validation is required.
	RequireEmailValidation bool `bson:"require_email_validation"`

	// LogErrors indicates whether to log errors during the authentication process.
	LogErrors bool `bson:"log_errors"`

	// AllowCustomersToRemoveAssociations indicates whether users can remove external authentication associations.
	AllowCustomersToRemoveAssociations bool `bson:"allow_customers_to_remove_associations"`

	// ActiveAuthenticationMethodSystemNames contains system names of active authentication methods.
	ActiveAuthenticationMethodSystemNames []string `bson:"active_authentication_method_system_names"`
}

// NewExternalAuthenticationSettings creates a new instance of ExternalAuthenticationSettings with default values.
func NewExternalAuthenticationSettings() *ExternalAuthenticationSettings {
	return &ExternalAuthenticationSettings{
		ActiveAuthenticationMethodSystemNames: []string{},
	}
}
