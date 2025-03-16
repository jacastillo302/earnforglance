package domain

// OrderVoidedEvent represents an order voided event
type OrderVoidedEvent struct {
	Order Order
}

// NewOrderVoidedEvent creates a new instance of OrderVoidedEvent
func NewOrderVoidedEvent(order Order) *OrderVoidedEvent {
	return &OrderVoidedEvent{
		Order: order,
	}
}
