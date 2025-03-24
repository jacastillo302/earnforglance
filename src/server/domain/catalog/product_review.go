package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductReview = "product_reviews"
)

// ProductReview represents a product review
type ProductReview struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID              primitive.ObjectID `bson:"customer_id"`
	ProductID               primitive.ObjectID `bson:"product_id"`
	StoreID                 primitive.ObjectID `bson:"store_id"`
	IsApproved              bool               `bson:"is_approved"`
	Title                   string             `bson:"title"`
	ReviewText              string             `bson:"review_text"`
	ReplyText               string             `bson:"reply_text"`
	CustomerNotifiedOfReply bool               `bson:"customer_notified_of_reply"`
	Rating                  int                `bson:"rating"`
	HelpfulYesTotal         int                `bson:"helpful_yes_total"`
	HelpfulNoTotal          int                `bson:"helpful_no_total"`
	CreatedOnUtc            time.Time          `bson:"created_on_utc"`
}

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
