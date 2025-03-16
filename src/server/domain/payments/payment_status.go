package domain

// PaymentStatus represents a payment status enumeration
type PaymentStatus int

const (
	// Pending represents a pending payment status
	Pending PaymentStatus = 10

	// Authorized represents an authorized payment status
	Authorized PaymentStatus = 20

	// Paid represents a paid payment status
	Paid PaymentStatus = 30

	// PartiallyRefunded represents a partially refunded payment status
	PartiallyRefunded PaymentStatus = 35

	// Refunded represents a refunded payment status
	Refunded PaymentStatus = 40

	// Voided represents a voided payment status
	Voided PaymentStatus = 50
)
