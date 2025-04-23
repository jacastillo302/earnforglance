package domain

import (
	"context" // Added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionVendorNote = "vendor_notes"
)

// VendorNote represents a vendor note
type VendorNote struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	VendorID     bson.ObjectID `bson:"vendor_id"`
	Note         string        `bson:"note"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

// VendorNoteRepository defines the repository interface for VendorNote
type VendorNoteRepository interface {
	CreateMany(c context.Context, items []VendorNote) error
	Create(c context.Context, vendor_note *VendorNote) error
	Update(c context.Context, vendor_note *VendorNote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorNote, error)
	FetchByID(c context.Context, ID string) (VendorNote, error)
}

// VendorNoteUsecase defines the use case interface for VendorNote
type VendorNoteUsecase interface {
	CreateMany(c context.Context, items []VendorNote) error
	FetchByID(c context.Context, ID string) (VendorNote, error)
	Create(c context.Context, vendor_note *VendorNote) error
	Update(c context.Context, vendor_note *VendorNote) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorNote, error)
}
