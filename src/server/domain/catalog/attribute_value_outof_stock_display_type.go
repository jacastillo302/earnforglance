package domain

// AttributeValueOutOfStockDisplayType represents an attribute value display type when out of stock
type AttributeValueOutOfStockDisplayType bool

const (
	// Disable represents an attribute value that is visible but cannot be interacted with
	Disable AttributeValueOutOfStockDisplayType = true

	// AlwaysDisplay represents an attribute value that is always displayed
	AlwaysDisplay AttributeValueOutOfStockDisplayType = false
)
