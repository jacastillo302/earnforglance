package domain

// CustomerChangeWorkingLanguageEvent represents a customer change working language event
type CustomerChangeWorkingLanguageEvent struct {
	Customer Customer
}

// NewCustomerChangeWorkingLanguageEvent creates a new CustomerChangeWorkingLanguageEvent
func NewCustomerChangeWorkingLanguageEvent(customer Customer) *CustomerChangeWorkingLanguageEvent {
	return &CustomerChangeWorkingLanguageEvent{
		Customer: customer,
	}
}
