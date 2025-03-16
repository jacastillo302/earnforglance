package domain

// EntityTokensAddedEvent is a container for tokens that are added
type EntityTokensAddedEvent[T any, U any] struct {
	Entity T
	Tokens []U
}

// NewEntityTokensAddedEvent creates a new instance of EntityTokensAddedEvent
func NewEntityTokensAddedEvent[T any, U any](entity T, tokens []U) *EntityTokensAddedEvent[T, U] {
	return &EntityTokensAddedEvent[T, U]{
		Entity: entity,
		Tokens: tokens,
	}
}
