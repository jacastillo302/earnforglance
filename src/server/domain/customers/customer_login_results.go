package domain

// CustomerLoginResults represents the result of a customer login attempt
type CustomerLoginResults int

const (
	// Successful indicates a successful login
	Successful CustomerLoginResults = 1

	// CustomerNotExist indicates that the customer does not exist (email or username)
	CustomerNotExist CustomerLoginResults = 2

	// WrongPassword indicates that the provided password is incorrect
	WrongPassword CustomerLoginResults = 3

	// NotActive indicates that the account has not been activated
	NotActive CustomerLoginResults = 4

	// Deleted indicates that the customer account has been deleted
	Deleted CustomerLoginResults = 5

	// NotRegistered indicates that the customer is not registered
	NotRegistered CustomerLoginResults = 6

	// LockedOut indicates that the customer account is locked out
	LockedOut CustomerLoginResults = 7

	// MultiFactorAuthenticationRequired indicates that multi-factor authentication is required
	MultiFactorAuthenticationRequired CustomerLoginResults = 8
)
