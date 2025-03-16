package domain

// OrderStatusChangedEvent represents an order status changed event
type OrderStatusChangedEvent struct {
	Order               Order
	PreviousOrderStatus OrderStatus
}

// NewOrderStatusChangedEvent creates a new instance of OrderStatusChangedEvent
func NewOrderStatusChangedEvent(order Order, previousOrderStatus OrderStatus) *OrderStatusChangedEvent {
	return &OrderStatusChangedEvent{
		Order:               order,
		PreviousOrderStatus: previousOrderStatus,
	}
}
