package domain

// VatNumberStatus represents the VAT number status enumeration
type VatNumberStatus int

const (
	// Unknown represents an unknown VAT number status
	Unknown VatNumberStatus = 0

	// Empty represents an empty VAT number status
	Empty VatNumberStatus = 10

	// Valid represents a valid VAT number status
	Valid VatNumberStatus = 20

	// Invalid represents an invalid VAT number status
	Invalid VatNumberStatus = 30
)
