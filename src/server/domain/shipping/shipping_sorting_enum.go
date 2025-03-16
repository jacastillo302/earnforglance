package domain

// ShippingSortingEnum represents the shipping methods' sorting
type ShippingSortingEnum int

const (
	// Position represents sorting by position (display order)
	Position ShippingSortingEnum = 1

	// ShippingCost represents sorting by shipping cost
	ShippingCost ShippingSortingEnum = 2
)
