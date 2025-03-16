package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductReviewHelpfulness = "product_review_helpfulness"
)

// ProductReviewHelpfulness represents a product review helpfulness
type ProductReviewHelpfulness struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductReviewID int                `bson:"product_review_id"`
	WasHelpful      bool               `bson:"was_helpful"`
	CustomerID      int                `bson:"customer_id"`
}
