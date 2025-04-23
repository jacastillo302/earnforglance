package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionReturnRequest = "return_requests"
)

// ReturnRequest represents a return request
type ReturnRequest struct {
	ID                    bson.ObjectID `bson:"_id,omitempty"`
	CustomNumber          string        `bson:"custom_number"`
	StoreID               bson.ObjectID `bson:"store_id"`
	OrderItemID           bson.ObjectID `bson:"order_item_id"`
	CustomerID            bson.ObjectID `bson:"customer_id"`
	Quantity              int           `bson:"quantity"`
	ReturnedQuantity      int           `bson:"returned_quantity"`
	ReasonForReturn       string        `bson:"reason_for_return"`
	RequestedAction       string        `bson:"requested_action"`
	CustomerComments      string        `bson:"customer_comments"`
	UploadedFileID        bson.ObjectID `bson:"uploaded_file_id"`
	StaffNotes            string        `bson:"staff_notes"`
	ReturnRequestStatusID int           `bson:"return_request_status_id"`
	CreatedOnUtc          time.Time     `bson:"created_on_utc"`
	UpdatedOnUtc          time.Time     `bson:"updated_on_utc"`
}

// ReturnRequestRepository represents the repository interface for ReturnRequest
type ReturnRequestRepository interface {
	CreateMany(c context.Context, items []ReturnRequest) error
	Create(c context.Context, return_request *ReturnRequest) error
	Update(c context.Context, return_request *ReturnRequest) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReturnRequest, error)
	FetchByID(c context.Context, ID string) (ReturnRequest, error)
}

// ReturnRequestUsecase represents the usecase interface for ReturnRequest
type ReturnRequestUsecase interface {
	CreateMany(c context.Context, items []ReturnRequest) error
	FetchByID(c context.Context, ID string) (ReturnRequest, error)
	Create(c context.Context, return_request *ReturnRequest) error
	Update(c context.Context, return_request *ReturnRequest) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ReturnRequest, error)
}
