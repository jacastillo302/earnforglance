package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionVendor = "vendors"
)

// Vendor represents a vendor
type Vendor struct {
	ID                             bson.ObjectID  `bson:"_id,omitempty"`
	Name                           string         `bson:"name"`
	Email                          string         `bson:"email"`
	Description                    string         `bson:"description"`
	PictureID                      bson.ObjectID  `bson:"picture_id"`
	AddressID                      bson.ObjectID  `bson:"address_id"`
	AdminComment                   string         `bson:"admin_comment"`
	Active                         bool           `bson:"active"`
	Deleted                        bool           `bson:"deleted"`
	DisplayOrder                   int            `bson:"display_order"`
	MetaKeywords                   string         `bson:"meta_keywords"`
	MetaDescription                string         `bson:"meta_description"`
	MetaTitle                      string         `bson:"meta_title"`
	PageSize                       int            `bson:"page_size"`
	AllowCustomersToSelectPageSize bool           `bson:"allow_customers_to_select_page_size"`
	PageSizeOptions                string         `bson:"page_size_options"`
	PriceRangeFiltering            bool           `bson:"price_range_filtering"`
	PriceFrom                      float64        `bson:"price_from"`
	PriceTo                        float64        `bson:"price_to"`
	ManuallyPriceRange             bool           `bson:"manually_price_range"`
	PmCustomerID                   *bson.ObjectID `bson:"pm_customer_id"`
}

// VendorRepository defines the repository interface for Vendor
type VendorRepository interface {
	CreateMany(c context.Context, items []Vendor) error
	Create(c context.Context, vendor *Vendor) error
	Update(c context.Context, vendor *Vendor) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Vendor, error)
	FetchByID(c context.Context, ID string) (Vendor, error)
}

// VendorUsecase defines the use case interface for Vendor
type VendorUsecase interface {
	CreateMany(c context.Context, items []Vendor) error
	FetchByID(c context.Context, ID string) (Vendor, error)
	Create(c context.Context, vendor *Vendor) error
	Update(c context.Context, vendor *Vendor) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Vendor, error)
}
