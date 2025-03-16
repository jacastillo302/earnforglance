package domain

// ForumSearchType represents a forum search type
type ForumSearchType int

const (
	// All represents searching in topic titles and post text
	All ForumSearchType = 0

	// TopicTitlesOnly represents searching in topic titles only
	TopicTitlesOnly ForumSearchType = 10

	// PostTextOnly represents searching in post text only
	PostTextOnly ForumSearchType = 20
)
