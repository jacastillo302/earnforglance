package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectioVendorNote = "vendor_notes"
)

// VendorNote represents a vendor note
type VendorNote struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	VendorID     int                `bson:"vendor_id"`
	Note         string             `bson:"note"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
