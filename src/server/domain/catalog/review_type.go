package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionReviewType = "review_types"
)

// ReviewType represents a review type
type ReviewType struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Name                  string             `bson:"name"`
	Description           string             `bson:"description"`
	DisplayOrder          int                `bson:"display_order"`
	VisibleToAllCustomers bool               `bson:"visible_to_all_customers"`
	IsRequired            bool               `bson:"is_required"`
}
