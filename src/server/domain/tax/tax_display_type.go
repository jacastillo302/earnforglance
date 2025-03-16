package domain

// TaxDisplayType represents the tax display type enumeration
type TaxDisplayType int

const (
	// IncludingTax represents including tax
	IncludingTax TaxDisplayType = 0

	// ExcludingTax represents excluding tax
	ExcludingTax TaxDisplayType = 10
)
