package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSecuritySettings = "security_settings"
)

// SecuritySettings represents security settings.
type SecuritySettings struct {
	ID                                                     primitive.ObjectID `bson:"_id,omitempty"`
	EncryptionKey                                          string             `bson:"encryption_key"`
	AdminAreaAllowedIpAddresses                            []string           `bson:"admin_area_allowed_ip_addresses"`
	HoneypotEnabled                                        bool               `bson:"honeypot_enabled"`
	HoneypotInputName                                      string             `bson:"honeypot_input_name"`
	LogHoneypotDetection                                   bool               `bson:"log_honeypot_detection"`
	AllowNonAsciiCharactersInHeaders                       bool               `bson:"allow_non_ascii_characters_in_headers"`
	UseAesEncryptionAlgorithm                              bool               `bson:"use_aes_encryption_algorithm"`
	AllowStoreOwnerExportImportCustomersWithHashedPassword bool               `bson:"allow_store_owner_export_import_customers_with_hashed_password"`
}

// NewSecuritySettings creates a new instance of SecuritySettings with default values
func NewSecuritySettings() *SecuritySettings {
	return &SecuritySettings{
		AdminAreaAllowedIpAddresses: []string{},
	}
}

type SecuritySettingsRepository interface {
	CreateMany(c context.Context, items []SecuritySettings) error
	Create(c context.Context, security_settings *SecuritySettings) error
	Update(c context.Context, security_settings *SecuritySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SecuritySettings, error)
	FetchByID(c context.Context, security_settingsID string) (SecuritySettings, error)
}

type SecuritySettingsUsecase interface {
	FetchByID(c context.Context, security_settingsID string) (SecuritySettings, error)
	Create(c context.Context, security_settings *SecuritySettings) error
	Update(c context.Context, security_settings *SecuritySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SecuritySettings, error)
}
