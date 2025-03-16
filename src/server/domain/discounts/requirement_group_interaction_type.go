package domain

// RequirementGroupInteractionType represents an interaction type within the group of requirements
type RequirementGroupInteractionType int

const (
	// And represents that all requirements within the group must be met
	And RequirementGroupInteractionType = 0

	// Or represents that at least one of the requirements within the group must be met
	Or RequirementGroupInteractionType = 2
)
