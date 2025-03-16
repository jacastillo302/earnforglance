package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionVendorAttributeValue = "vendor_attribute_values"
)

// VendorAttribute represents a vendor attribute
type VendorAttributeValue struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Add other fields from BaseAttribute if needed
}
