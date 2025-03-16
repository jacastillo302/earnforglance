package domain

// PasswordFormat represents the password format
type PasswordFormat int

const (
	// Clear represents a clear password format
	Clear PasswordFormat = 0

	// Hashed represents a hashed password format
	Hashed PasswordFormat = 1

	// Encrypted represents an encrypted password format
	Encrypted PasswordFormat = 2
)
