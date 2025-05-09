package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionVendorAttributeValue = "vendor_attribute_values"
)

// VendorAttribute represents a vendor attribute.
type VendorAttributeValue struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	// Add other fields from BaseAttribute if needed
}

type VendorAttributeValueRepository interface {
	CreateMany(c context.Context, items []VendorAttributeValue) error
	Create(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Update(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorAttributeValue, error)
	FetchByID(c context.Context, ID string) (VendorAttributeValue, error)
}

type VendorAttributeValueUsecase interface {
	CreateMany(c context.Context, items []VendorAttributeValue) error
	FetchByID(c context.Context, ID string) (VendorAttributeValue, error)
	Create(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Update(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorAttributeValue, error)
}
