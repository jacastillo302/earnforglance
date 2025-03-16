package domain

// SecuritySettingsChangedEvent represents a security setting changed event
type SecuritySettingsChangedEvent struct {
	SecuritySettings        SecuritySettings
	OldEncryptionPrivateKey string
}

// NewSecuritySettingsChangedEvent creates a new instance of SecuritySettingsChangedEvent
func NewSecuritySettingsChangedEvent(securitySettings SecuritySettings, oldEncryptionPrivateKey string) *SecuritySettingsChangedEvent {
	return &SecuritySettingsChangedEvent{
		SecuritySettings:        securitySettings,
		OldEncryptionPrivateKey: oldEncryptionPrivateKey,
	}
}
