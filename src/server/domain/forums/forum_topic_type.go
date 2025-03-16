package domain

// ForumTopicType represents a forum topic type
type ForumTopicType int

const (
	// Normal represents a normal forum topic
	Normal ForumTopicType = 10

	// Sticky represents a sticky forum topic
	Sticky ForumTopicType = 15

	// Announcement represents an announcement forum topic
	Announcement ForumTopicType = 20
)
