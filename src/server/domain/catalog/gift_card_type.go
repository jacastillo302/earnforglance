package domain

// GiftCardType represents a gift card type
type GiftCardType int

const (
	// Virtual represents a virtual gift card
	Virtual GiftCardType = 0

	// Physical represents a physical gift card
	Physical GiftCardType = 1
)
