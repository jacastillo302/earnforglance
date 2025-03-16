package domain

// EmailUnsubscribedEvent represents an email unsubscribed event
type EmailUnsubscribedEvent struct {
	Subscription *NewsLetterSubscription
}

// NewEmailUnsubscribedEvent creates a new instance of EmailUnsubscribedEvent
func NewEmailUnsubscribedEvent(subscription *NewsLetterSubscription) *EmailUnsubscribedEvent {
	return &EmailUnsubscribedEvent{
		Subscription: subscription,
	}
}

// Equals compares this instance to a specified EmailUnsubscribedEvent and returns an indication
func (e *EmailUnsubscribedEvent) Equals(other *EmailUnsubscribedEvent) bool {
	if other == nil {
		return false
	}
	if e == other {
		return true
	}
	return e.Subscription == other.Subscription
}

// Equals compares this instance to a specified object and returns an indication
func (e *EmailUnsubscribedEvent) EqualsObj(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if e == obj {
		return true
	}
	other, ok := obj.(*EmailUnsubscribedEvent)
	if !ok {
		return false
	}
	return e.Equals(other)
}
