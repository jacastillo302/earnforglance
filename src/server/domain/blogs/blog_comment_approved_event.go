package domain

// BlogCommentApprovedEvent represents a blog post comment approved event
type BlogCommentApprovedEvent struct {
	BlogComment BlogComment
}

// NewBlogCommentApprovedEvent creates a new BlogCommentApprovedEvent
func NewBlogCommentApprovedEvent(blogComment BlogComment) *BlogCommentApprovedEvent {
	return &BlogCommentApprovedEvent{
		BlogComment: blogComment,
	}
}
