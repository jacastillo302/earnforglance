package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionReturnRequestReason = "return_request_reasons"
)

// ReturnRequestReason represents a return request reason.
type ReturnRequestReason struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}

// ReturnRequestReasonRepository represents the repository interface for ReturnRequestReason
type ReturnRequestReasonRepository interface {
	Create(c context.Context, return_request_reason *ReturnRequestReason) error
	Update(c context.Context, return_request_reason *ReturnRequestReason) error
	Delete(c context.Context, return_request_reason *ReturnRequestReason) error
	Fetch(c context.Context) ([]ReturnRequestReason, error)
	FetchByID(c context.Context, return_request_reasonID string) (ReturnRequestReason, error)
}

// ReturnRequestReasonUsecase represents the use case interface for ReturnRequestReason
type ReturnRequestReasonUsecase interface {
	FetchByID(c context.Context, return_request_reasonID string) (ReturnRequestReason, error)
	Create(c context.Context, return_request_reason *ReturnRequestReason) error
	Update(c context.Context, return_request_reason *ReturnRequestReason) error
	Delete(c context.Context, return_request_reason *ReturnRequestReason) error
	Fetch(c context.Context) ([]ReturnRequestReason, error)
}
