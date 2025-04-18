package domain

import (
	address "earnforglance/server/domain/common"
	media "earnforglance/server/domain/media"
	domain "earnforglance/server/domain/vendors"
)

type VendorAddRequest struct {
	Vendor     domain.Vendor
	Attributes map[string]string
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
