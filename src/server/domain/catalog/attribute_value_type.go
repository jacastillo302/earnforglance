package domain

// AttributeValueType represents an attribute value type
type AttributeValueType int

const (
	// Simple represents a simple attribute value
	Simple AttributeValueType = 0

	// AssociatedToProduct represents an attribute value associated to a product (used when configuring bundled products)
	AssociatedToProduct AttributeValueType = 10
)
