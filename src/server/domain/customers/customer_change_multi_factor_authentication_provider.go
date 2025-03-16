package domain

// CustomerChangeMultiFactorAuthenticationProviderEvent represents a "Customer is change multi-factor authentication provider" event
type CustomerChangeMultiFactorAuthenticationProviderEvent struct {
	Customer Customer
}

// NewCustomerChangeMultiFactorAuthenticationProviderEvent creates a new CustomerChangeMultiFactorAuthenticationProviderEvent
func NewCustomerChangeMultiFactorAuthenticationProviderEvent(customer Customer) *CustomerChangeMultiFactorAuthenticationProviderEvent {
	return &CustomerChangeMultiFactorAuthenticationProviderEvent{
		Customer: customer,
	}
}
