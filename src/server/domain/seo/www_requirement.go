package domain

// WwwRequirement represents WWW requirement
type WwwRequirement int

const (
	// NoMatter means it doesn't matter (do nothing)
	NoMatter WwwRequirement = 0

	// WithWww means pages should have WWW prefix
	WithWww WwwRequirement = 10

	// WithoutWww means pages should not have WWW prefix
	WithoutWww WwwRequirement = 20
)
