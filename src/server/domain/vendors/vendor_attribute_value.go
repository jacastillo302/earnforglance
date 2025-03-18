package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionVendorAttributeValue = "vendor_attribute_values"
)

// VendorAttribute represents a vendor attribute.
type VendorAttributeValue struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Add other fields from BaseAttribute if needed
}

type VendorAttributeValueRepository interface {
	Create(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Update(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Delete(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Fetch(c context.Context) ([]VendorAttributeValue, error)
	FetchByID(c context.Context, vendor_attribute_valueID string) (VendorAttributeValue, error)
}

type VendorAttributeValueUsecase interface {
	FetchByID(c context.Context, vendor_attribute_valueID string) (VendorAttributeValue, error)
	Create(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Update(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Delete(c context.Context, vendor_attribute_value *VendorAttributeValue) error
	Fetch(c context.Context) ([]VendorAttributeValue, error)
}
