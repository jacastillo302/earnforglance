package domain

// AttributeValueOutOfStockDisplayType represents an attribute value display type when out of stock
type AttributeValueOutOfStockDisplayType int

const (
	// Disable represents an attribute value that is visible but cannot be interacted with
	Disable AttributeValueOutOfStockDisplayType = 0

	// AlwaysDisplay represents an attribute value that is always displayed
	AlwaysDisplay bool = false
)
