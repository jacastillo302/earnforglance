package domain

// DiscountLimitationType represents a discount limitation type
type DiscountLimitationType int

const (
	// Unlimited represents no limitation
	Unlimited DiscountLimitationType = 0

	// NTimesOnly represents a limitation of N times only
	NTimesOnly DiscountLimitationType = 15

	// NTimesPerCustomer represents a limitation of N times per customer
	NTimesPerCustomer DiscountLimitationType = 25
)
