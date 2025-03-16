package domain

import (
	"strings"
)

// CustomerExtensions provides extension methods for Customer
type CustomerExtensions struct{}

// IsSearchEngineAccount checks if the customer is a search engine account
func (ce *CustomerExtensions) IsSearchEngineAccount(customer *Customer) bool {
	if customer == nil {
		panic("customer cannot be nil")
	}

	if !customer.IsSystemAccount || customer.SystemName == "" {
		return false
	}

	return strings.EqualFold(customer.SystemName, CustomerDefaults.SearchEngineCustomerName)
}

// IsBackgroundTaskAccount checks if the customer is a built-in record for background tasks
func (ce *CustomerExtensions) IsBackgroundTaskAccount(customer *Customer) bool {
	if customer == nil {
		panic("customer cannot be nil")
	}

	if !customer.IsSystemAccount || customer.SystemName == "" {
		return false
	}

	return strings.EqualFold(customer.SystemName, CustomerDefaults.BackgroundTaskCustomerName)
}
