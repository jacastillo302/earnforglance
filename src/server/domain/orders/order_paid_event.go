package domain

// OrderPaidEvent represents an order paid event
type OrderPaidEvent struct {
	Order Order
}

// NewOrderPaidEvent creates a new instance of OrderPaidEvent
func NewOrderPaidEvent(order Order) *OrderPaidEvent {
	return &OrderPaidEvent{
		Order: order,
	}
}
