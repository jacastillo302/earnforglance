package domain

// ShippingStatus represents the shipping status enumeration
type ShippingStatus int

const (
	// ShippingNotRequired represents a shipping not required status
	ShippingNotRequired ShippingStatus = 10

	// NotYetShipped represents a not yet shipped status
	NotYetShipped ShippingStatus = 20

	// PartiallyShipped represents a partially shipped status
	PartiallyShipped ShippingStatus = 25

	// Shipped represents a shipped status
	Shipped ShippingStatus = 30

	// Delivered represents a delivered status
	Delivered ShippingStatus = 40
)
