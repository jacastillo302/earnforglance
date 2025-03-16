package domain

import "fmt"

// MessageDelayPeriodExtensions provides extension methods for MessageDelayPeriod
type MessageDelayPeriodExtensions struct{}

// ToHours returns message delay in hours
func (e *MessageDelayPeriodExtensions) ToHours(period MessageDelayPeriod, value int) int {
	switch period {
	case Hours:
		return value
	case Days:
		return value * 24
	default:
		panic(fmt.Sprintf("ArgumentOutOfRangeException: %v", period))
	}
}
