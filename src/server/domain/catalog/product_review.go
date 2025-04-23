package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionProductReview = "product_reviews"
)

// ProductReview represents a product review
type ProductReview struct {
	ID                      bson.ObjectID `bson:"_id,omitempty"`
	CustomerID              bson.ObjectID `bson:"customer_id"`
	ProductID               bson.ObjectID `bson:"product_id"`
	StoreID                 bson.ObjectID `bson:"store_id"`
	IsApproved              bool          `bson:"is_approved"`
	Title                   string        `bson:"title"`
	ReviewText              string        `bson:"review_text"`
	ReplyText               string        `bson:"reply_text"`
	CustomerNotifiedOfReply bool          `bson:"customer_notified_of_reply"`
	Rating                  int           `bson:"rating"`
	HelpfulYesTotal         int           `bson:"helpful_yes_total"`
	HelpfulNoTotal          int           `bson:"helpful_no_total"`
	CreatedOnUtc            time.Time     `bson:"created_on_utc"`
}

// JSON example for ProductReview
// {
//   "_id": "60c72b2f9af1c2b9d8e8b456",
//   "customer_id": "60c72b2f9af1c2b9d8e8b123",
//   "product_id": "60c72b2f9af1c2b9d8e8b789",
//   "store_id": "60c72b2f9af1c2b9d8e8b321",
//   "is_approved": true,
//   "title": "Great Product!",
//   "review_text": "I really enjoyed using this product. Highly recommend!",
//   "reply_text": "Thank you for your feedback!",
//   "customer_notified_of_reply": true,
//   "rating": 5,
//   "helpful_yes_total": 10,
//   "helpful_no_total": 2,
//   "created_on_utc": "2023-10-01T12:34:56Z"
// }

type ProductReviewRepository interface {
	CreateMany(c context.Context, items []ProductReview) error
	Create(c context.Context, product_review *ProductReview) error
	Update(c context.Context, product_review *ProductReview) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductReview, error)
	FetchByID(c context.Context, ID string) (ProductReview, error)
}

type ProductReviewUsecase interface {
	CreateMany(c context.Context, items []ProductReview) error
	FetchByID(c context.Context, ID string) (ProductReview, error)
	Create(c context.Context, product_review *ProductReview) error
	Update(c context.Context, product_review *ProductReview) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductReview, error)
}
