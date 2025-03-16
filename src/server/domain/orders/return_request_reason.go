package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionReturnRequestReason = "return_request_reasons"
)

// ReturnRequestReason represents a return request reason
type ReturnRequestReason struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
