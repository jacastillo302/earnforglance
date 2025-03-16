package domain

// ProductReviewApprovedEvent represents a product review approved event
type ProductReviewApprovedEvent struct {
	ProductReview ProductReview
}

// NewProductReviewApprovedEvent creates a new ProductReviewApprovedEvent
func NewProductReviewApprovedEvent(productReview ProductReview) *ProductReviewApprovedEvent {
	return &ProductReviewApprovedEvent{
		ProductReview: productReview,
	}
}
