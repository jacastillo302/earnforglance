package domain

import (
	domain "earnforglance/server/domain/customers"
)

// ResetCheckoutDataEvent represents a reset checkout data event
type ResetCheckoutDataEvent struct {
	Customer domain.Customer
	StoreID  int
}

// NewResetCheckoutDataEvent creates a new instance of ResetCheckoutDataEvent
func NewResetCheckoutDataEvent(customer domain.Customer, storeID int) *ResetCheckoutDataEvent {
	return &ResetCheckoutDataEvent{
		Customer: customer,
		StoreID:  storeID,
	}
}
