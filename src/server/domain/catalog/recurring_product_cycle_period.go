package domain

// RecurringProductCyclePeriod represents a recurring product cycle period
type RecurringProductCyclePeriod int

const (
	// Days represents a cycle period in days
	Days RecurringProductCyclePeriod = 0

	// Weeks represents a cycle period in weeks
	Weeks RecurringProductCyclePeriod = 10

	// Months represents a cycle period in months
	Months RecurringProductCyclePeriod = 20

	// Years represents a cycle period in years
	Years RecurringProductCyclePeriod = 30
)
