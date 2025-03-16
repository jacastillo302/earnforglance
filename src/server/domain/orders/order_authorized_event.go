package domain

// OrderAuthorizedEvent represents an order authorized event
type OrderAuthorizedEvent struct {
	Order Order
}

// NewOrderAuthorizedEvent creates a new instance of OrderAuthorizedEvent
func NewOrderAuthorizedEvent(order Order) *OrderAuthorizedEvent {
	return &OrderAuthorizedEvent{
		Order: order,
	}
}
