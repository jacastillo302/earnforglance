package domain

import (
	domain "earnforglance/server/domain/common"
	directory "earnforglance/server/domain/directory"
)

type Directory struct {
	Country       directory.Country
	StateProvince directory.StateProvince
}

type AddressAddRequest struct {
	Address                 domain.Address
	CustomAddressAttributes map[string]string
}

type AddressResponse struct {
	Addresses               domain.Address
	Directory               Directory
	CustomAddressAttributes map[string]string
}

type AddressesResponse struct {
	Addresses []AddressResponse
}

type BillingAddressRequest struct {
	AddressAddRequest
	ShipToSameAddress bool
	VatNumber         string
}

type BillingAddressResponse struct {
	AddressResponse
	EuVatEnabled          bool
	EuVatEnabledForGuests bool
	VatNumber             string
}
