package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionReturnRequestAction = "return_request_actions"
)

// ReturnRequestAction represents a return request action
type ReturnRequestAction struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
