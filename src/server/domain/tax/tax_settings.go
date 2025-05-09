package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionTaxSettings = "tax_settings"
)

// TaxSettings represents tax settings
type TaxSettings struct {
	ID                                    bson.ObjectID `bson:"_id,omitempty"`
	TaxBasedOnID                          int           `bson:"tax_based_on"`
	TaxBasedOnPickupPointAddress          bool          `bson:"tax_based_on_pickup_point_address"`
	TaxDisplayTypeID                      int           `bson:"tax_display_type"`
	ActiveTaxProviderSystemName           string        `bson:"active_tax_provider_system_name"`
	DefaultTaxAddressID                   bson.ObjectID `bson:"default_tax_address_id"`
	DisplayTaxSuffix                      bool          `bson:"display_tax_suffix"`
	DisplayTaxRates                       bool          `bson:"display_tax_rates"`
	PricesIncludeTax                      bool          `bson:"prices_include_tax"`
	AutomaticallyDetectCountry            bool          `bson:"automatically_detect_country"`
	AllowCustomersToSelectTaxDisplayType  bool          `bson:"allow_customers_to_select_tax_display_type"`
	HideZeroTax                           bool          `bson:"hide_zero_tax"`
	HideTaxInOrderSummary                 bool          `bson:"hide_tax_in_order_summary"`
	ForceTaxExclusionFromOrderSubtotal    bool          `bson:"force_tax_exclusion_from_order_subtotal"`
	DefaultTaxCategoryID                  bson.ObjectID `bson:"default_tax_category_id"`
	ShippingIsTaxable                     bool          `bson:"shipping_is_taxable"`
	ShippingPriceIncludesTax              bool          `bson:"shipping_price_includes_tax"`
	ShippingTaxClassID                    bson.ObjectID `bson:"shipping_tax_class_id"`
	PaymentMethodAdditionalFeeIsTaxable   bool          `bson:"payment_method_additional_fee_is_taxable"`
	PaymentMethodAdditionalFeeIncludesTax bool          `bson:"payment_method_additional_fee_includes_tax"`
	PaymentMethodAdditionalFeeTaxClassID  bson.ObjectID `bson:"payment_method_additional_fee_tax_class_id"`
	EuVatEnabled                          bool          `bson:"eu_vat_enabled"`
	EuVatRequired                         bool          `bson:"eu_vat_required"`
	EuVatEnabledForGuests                 bool          `bson:"eu_vat_enabled_for_guests"`
	EuVatShopCountryID                    bson.ObjectID `bson:"eu_vat_shop_country_id"`
	EuVatAllowVatExemption                bool          `bson:"eu_vat_allow_vat_exemption"`
	EuVatUseWebService                    bool          `bson:"eu_vat_use_web_service"`
	EuVatAssumeValid                      bool          `bson:"eu_vat_assume_valid"`
	EuVatEmailAdminWhenNewVatSubmitted    bool          `bson:"eu_vat_email_admin_when_new_vat_submitted"`
	LogErrors                             bool          `bson:"log_errors"`
}

// NewTaxSettings creates a new instance of TaxSettings with default values
func NewTaxSettings() *TaxSettings {
	return &TaxSettings{}
}

type TaxSettingsRepository interface {
	CreateMany(c context.Context, items []TaxSettings) error
	Create(c context.Context, tax_settings *TaxSettings) error
	Update(c context.Context, tax_settings *TaxSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]TaxSettings, error)
	FetchByID(c context.Context, ID string) (TaxSettings, error)
}

type TaxSettingsUsecase interface {
	CreateMany(c context.Context, items []TaxSettings) error
	FetchByID(c context.Context, ID string) (TaxSettings, error)
	Create(c context.Context, tax_settings *TaxSettings) error
	Update(c context.Context, tax_settings *TaxSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]TaxSettings, error)
}
