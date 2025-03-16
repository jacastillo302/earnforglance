package domain

// PostCopyProductEvent represents a post copy product event
type PostCopyProductEvent struct {
	OriginalProduct Product
	CopyProduct     Product
}

// NewPostCopyProductEvent creates a new PostCopyProductEvent
func NewPostCopyProductEvent(originalProduct Product, copyProduct Product) *PostCopyProductEvent {
	return &PostCopyProductEvent{
		OriginalProduct: originalProduct,
		CopyProduct:     copyProduct,
	}
}
