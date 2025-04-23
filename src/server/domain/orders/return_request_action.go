package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionReturnRequestAction = "return_request_actions"
)

// ReturnRequestAction represents a return request action.
type ReturnRequestAction struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	DisplayOrder int           `bson:"display_order"`
}

// ReturnRequestActionRepository interface
type ReturnRequestActionRepository interface {
	CreateMany(c context.Context, items []ReturnRequestAction) error
	Create(c context.Context, return_request_action *ReturnRequestAction) error
	Update(c context.Context, return_request_action *ReturnRequestAction) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReturnRequestAction, error)
	FetchByID(c context.Context, ID string) (ReturnRequestAction, error)
}

// ReturnRequestActionUsecase interface
type ReturnRequestActionUsecase interface {
	CreateMany(c context.Context, items []ReturnRequestAction) error
	FetchByID(c context.Context, ID string) (ReturnRequestAction, error)
	Create(c context.Context, return_request_action *ReturnRequestAction) error
	Update(c context.Context, return_request_action *ReturnRequestAction) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReturnRequestAction, error)
}
