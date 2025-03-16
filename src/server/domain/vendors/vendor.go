package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionVendor = "vendors"
)

// Vendor represents a vendor
type Vendor struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	Name                           string             `bson:"name"`
	Email                          string             `bson:"email"`
	Description                    string             `bson:"description"`
	PictureID                      int                `bson:"picture_id"`
	AddressID                      int                `bson:"address_id"`
	AdminComment                   string             `bson:"admin_comment"`
	Active                         bool               `bson:"active"`
	Deleted                        bool               `bson:"deleted"`
	DisplayOrder                   int                `bson:"display_order"`
	MetaKeywords                   string             `bson:"meta_keywords"`
	MetaDescription                string             `bson:"meta_description"`
	MetaTitle                      string             `bson:"meta_title"`
	PageSize                       int                `bson:"page_size"`
	AllowCustomersToSelectPageSize bool               `bson:"allow_customers_to_select_page_size"`
	PageSizeOptions                string             `bson:"page_size_options"`
	PriceRangeFiltering            bool               `bson:"price_range_filtering"`
	PriceFrom                      float64            `bson:"price_from"`
	PriceTo                        float64            `bson:"price_to"`
	ManuallyPriceRange             bool               `bson:"manually_price_range"`
	PmCustomerID                   *int               `bson:"pm_customer_id,omitempty"`
}
