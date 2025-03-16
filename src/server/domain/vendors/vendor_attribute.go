package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionVendorAttribute = "vendor_attributes"
)

// VendorAttribute represents a vendor attribute
type VendorAttribute struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Add other fields from BaseAttribute if needed
}
