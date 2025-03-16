package domain

// RentalPricePeriod represents a rental product period (for prices)
type RentalPricePeriod int

const (
	// Days represents a rental period in days
	RentalPricePeriod_Days RentalPricePeriod = 0

	// Weeks represents a rental period in weeks
	RentalPricePeriod_Weeks RentalPricePeriod = 10

	// Months represents a rental period in months
	RentalPricePeriod_Months RentalPricePeriod = 20

	// Years represents a rental period in years
	RentalPricePeriod_Years RentalPricePeriod = 30
)
