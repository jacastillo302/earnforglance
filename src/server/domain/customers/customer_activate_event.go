package domain

// CustomerActivatedEvent represents a customer activated event
type CustomerActivatedEvent struct {
	Customer Customer
}

// NewCustomerActivatedEvent creates a new CustomerActivatedEvent
func NewCustomerActivatedEvent(customer Customer) *CustomerActivatedEvent {
	return &CustomerActivatedEvent{
		Customer: customer,
	}
}
