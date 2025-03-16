package domain

// TaxBasedOn represents the tax based on
type TaxBasedOn int

const (
	// BillingAddress represents tax based on billing address
	BillingAddress TaxBasedOn = 1

	// ShippingAddress represents tax based on shipping address
	ShippingAddress TaxBasedOn = 2

	// DefaultAddress represents tax based on default address
	DefaultAddress TaxBasedOn = 3
)
