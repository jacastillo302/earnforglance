package domain

// ProductSortingEnum represents the product sorting
type ProductSortingEnum int

const (
	// Position represents sorting by position (display order)
	Position ProductSortingEnum = 0

	// NameAsc represents sorting by name: A to Z
	NameAsc ProductSortingEnum = 5

	// NameDesc represents sorting by name: Z to A
	NameDesc ProductSortingEnum = 6

	// PriceAsc represents sorting by price: Low to High
	PriceAsc ProductSortingEnum = 10

	// PriceDesc represents sorting by price: High to Low
	PriceDesc ProductSortingEnum = 11

	// CreatedOn represents sorting by product creation date
	CreatedOn ProductSortingEnum = 15
)
