package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductReviewHelpfulness = "product_review_helpfulness"
)

// ProductReviewHelpfulness represents a product review helpfulness
type ProductReviewHelpfulness struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	ProductReviewID bson.ObjectID `bson:"product_review_id"`
	WasHelpful      bool          `bson:"was_helpful"`
	CustomerID      bson.ObjectID `bson:"customer_id"`
}

type ProductReviewHelpfulnessRepository interface {
	CreateMany(c context.Context, items []ProductReviewHelpfulness) error
	Create(c context.Context, product_review_helpfulness *ProductReviewHelpfulness) error
	Update(c context.Context, product_review_helpfulness *ProductReviewHelpfulness) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductReviewHelpfulness, error)
	FetchByID(c context.Context, ID string) (ProductReviewHelpfulness, error)
}

type ProductReviewHelpfulnessUsecase interface {
	CreateMany(c context.Context, items []ProductReviewHelpfulness) error
	FetchByID(c context.Context, ID string) (ProductReviewHelpfulness, error)
	Create(c context.Context, product_review_helpfulnes *ProductReviewHelpfulness) error
	Update(c context.Context, product_review_helpfulness *ProductReviewHelpfulness) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductReviewHelpfulness, error)
}
