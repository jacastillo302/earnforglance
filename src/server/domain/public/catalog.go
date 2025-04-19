package domain

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	media "earnforglance/server/domain/media"
	shipping "earnforglance/server/domain/shipping"
	tax "earnforglance/server/domain/tax"
	vendor "earnforglance/server/domain/vendors"
)

type Filter struct {
	Field    string
	Value    string
	Operator string
}

type ProductRequest struct {
	ID                       string
	Filters                  []Filter
	Sort                     string
	Limit                    int
	Lang                     string
	ShowOnHomepage           bool
	IsRental                 bool
	IsTaxExempt              bool
	DisplayStockAvailability bool
	MarkAsNew                bool
	MinPrice                 float64
	MxnPrice                 float64
	Categories               []string
	Content                  []string
}

type ProductResponse struct {
	Product        domain.Product
	Template       domain.ProductTemplate
	Categories     []domain.Category
	Specifications []SpecificationAttribute
	Attributes     []ProductAttribute
	Pictures       []media.Picture
	Manufacturers  []domain.Manufacturer
	DeliveryDate   shipping.DeliveryDate
	Range          shipping.ProductAvailabilityRange
	Warehouse      Warehouse
	Tax            tax.TaxCategory
	Vendor         vendor.Vendor
	Reviews        []ProductReview
	Tags           []domain.ProductTag
	Videos         []media.Video
	Relates        []domain.Product
	Cross          []domain.Product
	TierPrice      []domain.TierPrice
	Download       *media.Download
}

type ProductsResponse struct {
	Products []ProductResponse
}

type ProductAttributeCombination struct {
	Value    domain.ProductAttributeCombination
	Pictures []media.Picture
}

type ProductAttributeValue struct {
	Value    domain.ProductAttributeValue
	Pictures []media.Picture
}

type ProductAttribute struct {
	Attribute    domain.ProductAttribute
	Values       []ProductAttributeValue
	Combinations []ProductAttributeCombination
}

type ProductReview struct {
	Review      domain.ProductReview
	Type        string
	Customer    string
	Helpfulness []domain.ProductReviewHelpfulness
}

type SpecificationAttribute struct {
	Attribute domain.SpecificationAttribute
	Group     domain.SpecificationAttributeGroup
	Options   []domain.SpecificationAttributeOption
}

type Warehouse struct {
	Warehouse shipping.Warehouse
	Inventory domain.ProductWarehouseInventory
}

type CatalogRepository interface {
	GetProducts(c context.Context, filter ProductRequest) ([]ProductsResponse, error)
}

type CatalogtUsecase interface {
	GetProducts(c context.Context, filter ProductRequest) ([]ProductsResponse, error)
}
