package domain

// ProductType represents a product type
type ProductType int

const (
	// SimpleProduct represents a simple product
	SimpleProduct ProductType = 5

	// GroupedProduct represents a grouped product (product with variants)
	GroupedProduct ProductType = 10
)
