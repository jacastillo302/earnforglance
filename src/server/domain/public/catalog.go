package domain

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	directory "earnforglance/server/domain/directory"
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

type ManufacturerRequest struct {
	ID                  string
	Filters             []Filter
	Sort                string
	Limit               int
	Page                int
	Lang                string
	PriceRangeFiltering bool
	ManuallyPriceRange  bool
	PriceFrom           float64
	PriceTo             float64
	Content             []string
}

type CategoryRequest struct {
	ID                  string
	Filters             []Filter
	Sort                string
	Limit               int
	Page                int
	Lang                string
	ShowOnHomepage      bool
	IncludeInTopMenu    bool
	PriceRangeFiltering bool
	ManuallyPriceRange  bool
	PriceFrom           float64
	PriceTo             float64
	Parent              string
	Content             []string
}

type ProductRequest struct {
	ID                       string
	Filters                  []Filter
	Sort                     string
	Limit                    int
	Page                     int
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

type CategoryResponse struct {
	Category domain.Category
	Template domain.CategoryTemplate
	Picture  media.Picture
	Childs   []CategoryChilds
}

type ManufacturerResponse struct {
	Manufacturer domain.Manufacturer
	Template     domain.ManufacturerTemplate
	Picture      media.Picture
}

type CategoryChilds struct {
	Category domain.Category
	Picture  media.Picture
}

type ProductResponse struct {
	Product                         domain.Product
	Type                            Type
	Template                        domain.ProductTemplate
	Categories                      []domain.Category
	Specifications                  []SpecificationAttribute
	Attributes                      []ProductAttribute
	Pictures                        []media.Picture
	Manufacturers                   []domain.Manufacturer
	DeliveryDate                    shipping.DeliveryDate
	Range                           shipping.ProductAvailabilityRange
	Warehouse                       Warehouse
	Tax                             tax.TaxCategory
	Vendor                          vendor.Vendor
	Reviews                         []ProductReview
	Tags                            []domain.ProductTag
	Videos                          []media.Video
	Relates                         []domain.Product
	Cross                           []domain.Product
	TierPrice                       []domain.TierPrice
	BasepriceUnit                   *directory.MeasureWeight
	BasepriceBaseUnit               *directory.MeasureWeight
	Download                        *media.Download
	DownloadType                    *Type
	RecurringProductCyclePeriodType *Type
	RentalPricePeriodType           *Type
	LowStockActivityType            *Type
	BackorderModeType               *Type
}

type ProductsResponse struct {
	Products []ProductResponse
}

type CategoriesResponse struct {
	Categories []CategoryResponse
}

type ManufacturersResponse struct {
	Manufacturers []ManufacturerResponse
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

type Type struct {
	Name        string
	Value       int
	Description string
}

type CatalogRepository interface {
	GetCategories(c context.Context, filter CategoryRequest) ([]CategoriesResponse, error)
	GetProducts(c context.Context, filter ProductRequest) ([]ProductsResponse, error)
	GetManufacturers(c context.Context, filter ManufacturerRequest) ([]ManufacturersResponse, error)
}

type CatalogtUsecase interface {
	GetCategories(c context.Context, filter CategoryRequest) ([]CategoriesResponse, error)
	GetProducts(c context.Context, filter ProductRequest) ([]ProductsResponse, error)
	GetManufacturers(c context.Context, filter ManufacturerRequest) ([]ManufacturersResponse, error)
}
