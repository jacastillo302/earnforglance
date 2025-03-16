package domain

// OrderStatus represents an order status enumeration
type OrderStatus int

const (
	// Pending represents a pending order status
	Pending OrderStatus = 10

	// Processing represents a processing order status
	Processing OrderStatus = 20

	// Complete represents a complete order status
	Complete OrderStatus = 30

	// Cancelled represents a cancelled order status
	Cancelled OrderStatus = 40
)
