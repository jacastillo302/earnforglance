package domain

// BackorderMode represents a backorder mode
type BackorderMode int

const (
	// NoBackorders represents no backorders
	NoBackorders BackorderMode = 0

	// AllowQtyBelow0 represents allowing quantity below 0
	AllowQtyBelow0 BackorderMode = 1

	// AllowQtyBelow0AndNotifyCustomer represents allowing quantity below 0 and notifying the customer
	AllowQtyBelow0AndNotifyCustomer BackorderMode = 2
)
