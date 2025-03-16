package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProductReviewReviewTypeMapping = "product_review_review_type_mappings"
)

// ProductReviewReviewTypeMapping represents a product review and review type mapping
type ProductReviewReviewTypeMapping struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductReviewID int                `bson:"product_review_id"`
	ReviewTypeID    int                `bson:"review_type_id"`
	Rating          int                `bson:"rating"`
}
