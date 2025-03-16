package domain

// OrderPlacedEvent represents an order placed event
type OrderPlacedEvent struct {
	Order Order
}

// NewOrderPlacedEvent creates a new instance of OrderPlacedEvent
func NewOrderPlacedEvent(order Order) *OrderPlacedEvent {
	return &OrderPlacedEvent{
		Order: order,
	}
}
