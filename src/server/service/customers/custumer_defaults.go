package service

// CustomerServicesDefaults represents default values related to customer services
type CustomerServicesDefaults struct{}

// PasswordSaltKeySize returns the password salt key size
func (CustomerServicesDefaults) PasswordSaltKeySize() int {
	return 5
}

// CustomerUsernameLength returns the max username length
func (CustomerServicesDefaults) CustomerUsernameLength() int {
	return 100
}

// DefaultHashedPasswordFormat returns the default hash format for customer passwords
func (CustomerServicesDefaults) DefaultHashedPasswordFormat() string {
	return "SHA512"
}

// CustomerAttributePrefix returns the default prefix for customer attributes
func (CustomerServicesDefaults) CustomerAttributePrefix() string {
	return "customer_attribute_"
}

// CustomerDeletedSuffix returns the default suffix for deleted customer records
func (CustomerServicesDefaults) CustomerDeletedSuffix() string {
	return "-DELETED"
}
