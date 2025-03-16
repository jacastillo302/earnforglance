package domain

// QueuedEmailPriority represents priority of queued email
type QueuedEmailPriority int

const (
	// Low represents low priority
	Low QueuedEmailPriority = iota

	// High represents high priority
	High = 5
)
