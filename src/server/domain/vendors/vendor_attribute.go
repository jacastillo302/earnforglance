package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionVendorAttribute = "vendor_attributes"
)

// VendorAttribute represents a vendor attribute.
type VendorAttribute struct {
	ID bson.ObjectID `bson:"_id,omitempty"`
	// Add other fields from BaseAttribute if needed
}

type VendorAttributeRepository interface {
	CreateMany(c context.Context, items []VendorAttribute) error
	Create(c context.Context, vendor_attribute *VendorAttribute) error
	Update(c context.Context, vendor_attribute *VendorAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorAttribute, error)
	FetchByID(c context.Context, ID string) (VendorAttribute, error)
}

type VendorAttributeUsecase interface {
	CreateMany(c context.Context, items []VendorAttribute) error
	FetchByID(c context.Context, ID string) (VendorAttribute, error)
	Create(c context.Context, vendor_attribute *VendorAttribute) error
	Update(c context.Context, vendor_attribute *VendorAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorAttribute, error)
}
