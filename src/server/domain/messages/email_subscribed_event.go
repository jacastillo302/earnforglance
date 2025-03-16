package domain

// EmailSubscribedEvent represents an email subscribed event
type EmailSubscribedEvent struct {
	Subscription *NewsLetterSubscription
}

// NewEmailSubscribedEvent creates a new instance of EmailSubscribedEvent
func NewEmailSubscribedEvent(subscription *NewsLetterSubscription) *EmailSubscribedEvent {
	return &EmailSubscribedEvent{
		Subscription: subscription,
	}
}

// Equals compares this instance to a specified EmailSubscribedEvent and returns an indication
func (e *EmailSubscribedEvent) Equals(other *EmailSubscribedEvent) bool {
	if other == nil {
		return false
	}
	if e == other {
		return true
	}
	return e.Subscription == other.Subscription
}

// Equals compares this instance to a specified object and returns an indication
func (e *EmailSubscribedEvent) EqualsObj(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if e == obj {
		return true
	}
	other, ok := obj.(*EmailSubscribedEvent)
	if !ok {
		return false
	}
	return e.Equals(other)
}
