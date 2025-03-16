package domain

// OrderRefundedEvent represents an order refunded event
type OrderRefundedEvent struct {
	Order  Order
	Amount float64
}

// NewOrderRefundedEvent creates a new instance of OrderRefundedEvent
func NewOrderRefundedEvent(order Order, amount float64) *OrderRefundedEvent {
	return &OrderRefundedEvent{
		Order:  order,
		Amount: amount,
	}
}
