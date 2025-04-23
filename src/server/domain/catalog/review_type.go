package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionReviewType = "review_types"
)

// ReviewType represents a review type
type ReviewType struct {
	ID                    bson.ObjectID `bson:"_id,omitempty"`
	Name                  string        `bson:"name"`
	Description           string        `bson:"description"`
	DisplayOrder          int           `bson:"display_order"`
	VisibleToAllCustomers bool          `bson:"visible_to_all_customers"`
	IsRequired            bool          `bson:"is_required"`
}

type ReviewTypeRepository interface {
	CreateMany(c context.Context, items []ReviewType) error
	Create(c context.Context, review_type *ReviewType) error
	Update(c context.Context, review_type *ReviewType) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReviewType, error)
	FetchByID(c context.Context, ID string) (ReviewType, error)
}

type ReviewTypeUsecase interface {
	CreateMany(c context.Context, items []ReviewType) error
	FetchByID(c context.Context, ID string) (ReviewType, error)
	Create(c context.Context, review_type *ReviewType) error
	Update(c context.Context, review_type *ReviewType) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReviewType, error)
}
