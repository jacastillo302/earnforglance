package domain

// UserRegistrationType represents the customer registration type formatting enumeration
type UserRegistrationType int

const (
	// Standard represents standard account creation
	Standard UserRegistrationType = 1

	// EmailValidation represents email validation required after registration
	EmailValidation UserRegistrationType = 2

	// AdminApproval represents a customer should be approved by administrator
	AdminApproval UserRegistrationType = 3

	// Disabled represents registration is disabled
	Disabled UserRegistrationType = 4
)
