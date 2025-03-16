package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrderNote = "order_notes"
)

// OrderNote represents an order note
type OrderNote struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	OrderID           int                `bson:"order_id"`
	Note              string             `bson:"note"`
	DownloadID        int                `bson:"download_id"`
	DisplayToCustomer bool               `bson:"display_to_customer"`
	CreatedOnUtc      time.Time          `bson:"created_on_utc"`
}
