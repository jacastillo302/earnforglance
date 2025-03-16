package domain

// MessageTokensAddedEvent is a container for tokens that are added
type MessageTokensAddedEvent[U any] struct {
	Message MessageTemplate
	Tokens  []U
}

// NewMessageTokensAddedEvent creates a new instance of MessageTokensAddedEvent
func NewMessageTokensAddedEvent[U any](message MessageTemplate, tokens []U) *MessageTokensAddedEvent[U] {
	return &MessageTokensAddedEvent[U]{
		Message: message,
		Tokens:  tokens,
	}
}
