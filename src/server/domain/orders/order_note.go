package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionOrderNote = "order_notes"
)

// OrderNote represents an order note
type OrderNote struct {
	ID                bson.ObjectID `bson:"_id,omitempty"`
	OrderID           bson.ObjectID `bson:"order_id"`
	Note              string        `bson:"note"`
	DownloadID        bson.ObjectID `bson:"download_id"`
	DisplayToCustomer bool          `bson:"display_to_customer"`
	CreatedOnUtc      time.Time     `bson:"created_on_utc"`
}

// OrderNoteRepository interface
type OrderNoteRepository interface {
	CreateMany(c context.Context, items []OrderNote) error
	Create(c context.Context, order_note *OrderNote) error
	Update(c context.Context, order_note *OrderNote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]OrderNote, error)
	FetchByID(c context.Context, ID string) (OrderNote, error)
}

// OrderNoteUsecase interface
type OrderNoteUsecase interface {
	CreateMany(c context.Context, items []OrderNote) error
	FetchByID(c context.Context, ID string) (OrderNote, error)
	Create(c context.Context, order_note *OrderNote) error
	Update(c context.Context, order_note *OrderNote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]OrderNote, error)
}
