package domain

// LowStockActivity represents a low stock activity
type LowStockActivity int

const (
	// Nothing represents no action
	Nothing LowStockActivity = 0

	// DisableBuyButton represents disabling the buy button
	DisableBuyButton LowStockActivity = 1

	// Unpublish represents unpublishing the product
	Unpublish LowStockActivity = 2
)
