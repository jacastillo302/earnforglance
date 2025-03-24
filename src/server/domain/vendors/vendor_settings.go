package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionVendorSettings = "vendor_settings"
)

// VendorSettings represents vendor settings.
type VendorSettings struct {
	ID                                           primitive.ObjectID `bson:"_id,omitempty"`
	DefaultVendorPageSizeOptions                 string             `bson:"default_vendor_page_size_options"`
	VendorsBlockItemsToDisplay                   int                `bson:"vendors_block_items_to_display"`
	ShowVendorOnProductDetailsPage               bool               `bson:"show_vendor_on_product_details_page"`
	ShowVendorOnOrderDetailsPage                 bool               `bson:"show_vendor_on_order_details_page"`
	AllowCustomersToContactVendors               bool               `bson:"allow_customers_to_contact_vendors"`
	AllowCustomersToApplyForVendorAccount        bool               `bson:"allow_customers_to_apply_for_vendor_account"`
	TermsOfServiceEnabled                        bool               `bson:"terms_of_service_enabled"`
	AllowSearchByVendor                          bool               `bson:"allow_search_by_vendor"`
	AllowVendorsToEditInfo                       bool               `bson:"allow_vendors_to_edit_info"`
	NotifyStoreOwnerAboutVendorInformationChange bool               `bson:"notify_store_owner_about_vendor_information_change"`
	MaximumProductNumber                         int                `bson:"maximum_product_number"`
	AllowVendorsToImportProducts                 bool               `bson:"allow_vendors_to_import_products"`
	MaximumProductPicturesNumber                 int                `bson:"maximum_product_pictures_number"`
}

// NewVendorSettings creates a new instance of VendorSettings with default values
func NewVendorSettings() *VendorSettings {
	return &VendorSettings{}
}

// VendorSettingsRepository defines the repository interface for VendorSettings
type VendorSettingsRepository interface {
	CreateMany(c context.Context, items []VendorSettings) error
	Create(c context.Context, vendor_settings *VendorSettings) error
	Update(c context.Context, vendor_settings *VendorSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorSettings, error)
	FetchByID(c context.Context, ID string) (VendorSettings, error)
}

// VendorSettingsUsecase defines the use case interface for VendorSettings
type VendorSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (VendorSettings, error)
	Create(c context.Context, vendor_settings *VendorSettings) error
	Update(c context.Context, vendor_settings *VendorSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]VendorSettings, error)
}
