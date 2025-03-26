package domain

// RoundingType represents the rounding type
type RoundingType int

const (
	// Rounding001 represents default rounding (Match.Round(num, 2))
	Rounding001 RoundingType = 0

	// Rounding005Up represents prices rounded up to the nearest multiple of 5 cents
	Rounding005Up RoundingType = 10

	// Rounding005Down represents prices rounded down to the nearest multiple of 5 cents
	Rounding005Down RoundingType = 20

	// Rounding01Up represents rounding up to the nearest 10 cent value
	Rounding01Up RoundingType = 30

	// Rounding01Down represents rounding down to the nearest 10 cent value
	Rounding01Down RoundingType = 40

	// Rounding05 represents various rounding rules for sales ending in different cents
	Rounding05 RoundingType = 50

	// Rounding1 represents rounding rules for sales ending in different cents
	Rounding1 RoundingType = 60

	// Rounding1Up represents sales ending in 1–99 cents round up to the next whole dollar
	Rounding1Up RoundingType = 70
)
