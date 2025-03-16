package domain

// DiscountType represents a discount type
type DiscountType int

const (
	// AssignedToOrderTotal represents a discount assigned to order total
	AssignedToOrderTotal DiscountType = 1

	// AssignedToSkus represents a discount assigned to products (SKUs)
	AssignedToSkus DiscountType = 2

	// AssignedToCategories represents a discount assigned to categories (all products in a category)
	AssignedToCategories DiscountType = 5

	// AssignedToManufacturers represents a discount assigned to manufacturers (all products of a manufacturer)
	AssignedToManufacturers DiscountType = 6

	// AssignedToShipping represents a discount assigned to shipping
	AssignedToShipping DiscountType = 10

	// AssignedToOrderSubTotal represents a discount assigned to order subtotal
	AssignedToOrderSubTotal DiscountType = 20
)
