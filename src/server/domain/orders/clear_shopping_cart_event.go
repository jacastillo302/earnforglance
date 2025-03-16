package domain

// ClearShoppingCartEvent represents a shopping cart cleared event
type ClearShoppingCartEvent struct {
	Cart []ShoppingCartItem
}

// NewClearShoppingCartEvent creates a new instance of ClearShoppingCartEvent
func NewClearShoppingCartEvent(cart []ShoppingCartItem) *ClearShoppingCartEvent {
	return &ClearShoppingCartEvent{
		Cart: cart,
	}
}
