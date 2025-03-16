package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductReview = "product_reviews"
)

// ProductReview represents a product review
type ProductReview struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID              int                `bson:"customer_id"`
	ProductID               int                `bson:"product_id"`
	StoreID                 int                `bson:"store_id"`
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
