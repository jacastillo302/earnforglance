package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrderNote = "order_notes"
)

// OrderNote represents an order note
type OrderNote struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	OrderID           primitive.ObjectID `bson:"order_id"`
	Note              string             `bson:"note"`
	DownloadID        primitive.ObjectID `bson:"download_id"`
	DisplayToCustomer bool               `bson:"display_to_customer"`
	CreatedOnUtc      time.Time          `bson:"created_on_utc"`
}

// OrderNoteRepository interface
type OrderNoteRepository interface {
	Create(c context.Context, order_note *OrderNote) error
	Update(c context.Context, order_note *OrderNote) error
	Delete(c context.Context, order_note *OrderNote) error
	Fetch(c context.Context) ([]OrderNote, error)
	FetchByID(c context.Context, order_noteID string) (OrderNote, error)
}

// OrderNoteUsecase interface
type OrderNoteUsecase interface {
	FetchByID(c context.Context, order_noteID string) (OrderNote, error)
	Create(c context.Context, order_note *OrderNote) error
	Update(c context.Context, order_note *OrderNote) error
	Delete(c context.Context, order_note *OrderNote) error
	Fetch(c context.Context) ([]OrderNote, error)
}
