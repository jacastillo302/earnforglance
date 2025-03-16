package domain

// AdditionalTokensAddedEvent represents an event for "Additional tokens added"
type AdditionalTokensAddedEvent struct {
	AdditionalTokens []string
	TokenGroups      []string
}

// NewAdditionalTokensAddedEvent creates a new instance of AdditionalTokensAddedEvent
func NewAdditionalTokensAddedEvent() *AdditionalTokensAddedEvent {
	return &AdditionalTokensAddedEvent{
		AdditionalTokens: []string{},
	}
}

// AddTokens adds tokens to the AdditionalTokens list
func (e *AdditionalTokensAddedEvent) AddTokens(additionalTokens ...string) {
	e.AdditionalTokens = append(e.AdditionalTokens, additionalTokens...)
}
