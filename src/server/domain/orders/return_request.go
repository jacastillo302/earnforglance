package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionReturnRequest = "return_requests"
)

// ReturnRequest represents a return request
type ReturnRequest struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty"`
	CustomNumber          string              `bson:"custom_number"`
	StoreID               int                 `bson:"store_id"`
	OrderItemID           int                 `bson:"order_item_id"`
	CustomerID            int                 `bson:"customer_id"`
	Quantity              int                 `bson:"quantity"`
	ReturnedQuantity      int                 `bson:"returned_quantity"`
	ReasonForReturn       string              `bson:"reason_for_return"`
	RequestedAction       string              `bson:"requested_action"`
	CustomerComments      string              `bson:"customer_comments"`
	UploadedFileID        int                 `bson:"uploaded_file_id"`
	StaffNotes            string              `bson:"staff_notes"`
	ReturnRequestStatusID int                 `bson:"return_request_status_id"`
	CreatedOnUtc          time.Time           `bson:"created_on_utc"`
	UpdatedOnUtc          time.Time           `bson:"updated_on_utc"`
	ReturnRequestStatus   ReturnRequestStatus `bson:"return_request_status"`
}
