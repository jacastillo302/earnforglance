package domain

// ProductUrlStructureType represents the product URL structure type enum
type ProductUrlStructureType int

const (
	// Product represents the URL structure for product only (e.g. '/product-seo-name')
	ProductSeo ProductUrlStructureType = 0

	// CategoryProduct represents the URL structure for category (the most nested), then product (e.g. '/category-seo-name/product-seo-name')
	CategoryProduct ProductUrlStructureType = 10

	// ManufacturerProduct represents the URL structure for manufacturer, then product (e.g. '/manufacturer-seo-name/product-seo-name')
	ManufacturerProduct ProductUrlStructureType = 20
)
