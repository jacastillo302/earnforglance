package domain

import (
	"context"
	address "earnforglance/server/domain/common"
	media "earnforglance/server/domain/media"
	domain "earnforglance/server/domain/vendors"
)

type VendorRequest struct {
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

type VendorResponse struct {
	Vendor     domain.Vendor
	Address    address.Address
	Picture    *media.Picture
	Attributes map[string]string
}

type VendorsResponse struct {
	Vendors []VendorResponse
}

type VendorRepository interface {
	GetVendors(c context.Context, filter VendorRequest) ([]VendorsResponse, error)
}

type VendortUsecase interface {
	GetVendors(c context.Context, filter VendorRequest) ([]VendorsResponse, error)
}
