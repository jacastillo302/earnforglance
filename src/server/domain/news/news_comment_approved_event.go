package domain

// NewsCommentApprovedEvent represents a news comment approved event
type NewsCommentApprovedEvent struct {
	NewsComment NewsComment
}

// NewNewsCommentApprovedEvent creates a new instance of NewsCommentApprovedEvent
func NewNewsCommentApprovedEvent(newsComment NewsComment) *NewsCommentApprovedEvent {
	return &NewsCommentApprovedEvent{
		NewsComment: newsComment,
	}
}
