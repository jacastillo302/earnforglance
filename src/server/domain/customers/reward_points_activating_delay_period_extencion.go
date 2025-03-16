package domain

import "fmt"

// RewardPointsActivatingDelayPeriodExtensions provides extension methods for RewardPointsActivatingDelayPeriod
type RewardPointsActivatingDelayPeriodExtensions struct{}

// ToHours returns a delay period before activating points in hours
func (e *RewardPointsActivatingDelayPeriodExtensions) ToHours(period RewardPointsActivatingDelayPeriod, value int) int {
	switch period {
	case Hours:
		return value
	case Days:
		return value * 24
	default:
		panic(fmt.Sprintf("invalid RewardPointsActivatingDelayPeriod: %v", period))
	}
}
