package domain

import (
	address "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/shipping"
)

type WarehouseResponse struct {
	Warehouse  domain.Warehouse
	Address    address.Address
	Attributes map[string]string
}

type WarehousesResponse struct {
	Warehouses []WarehouseResponse
}
