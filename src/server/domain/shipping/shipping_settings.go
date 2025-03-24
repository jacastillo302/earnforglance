package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShippingSettings = "shipping_settings"
)

// ShippingSettings represents shipping settings.
type ShippingSettings struct {
	ID                                               primitive.ObjectID  `bson:"_id,omitempty"`
	ActiveShippingRateComputationMethodSystemNames   []string            `bson:"active_shipping_rate_computation_method_system_names"`
	ActivePickupPointProviderSystemNames             []string            `bson:"active_pickup_point_provider_system_names"`
	ShipToSameAddress                                bool                `bson:"ship_to_same_address"`
	AllowPickupInStore                               bool                `bson:"allow_pickup_in_store"`
	DisplayPickupPointsOnMap                         bool                `bson:"display_pickup_points_on_map"`
	IgnoreAdditionalShippingChargeForPickupInStore   bool                `bson:"ignore_additional_shipping_charge_for_pickup_in_store"`
	GoogleMapsApiKey                                 string              `bson:"google_maps_api_key"`
	UseWarehouseLocation                             bool                `bson:"use_warehouse_location"`
	NotifyCustomerAboutShippingFromMultipleLocations bool                `bson:"notify_customer_about_shipping_from_multiple_locations"`
	FreeShippingOverXEnabled                         bool                `bson:"free_shipping_over_x_enabled"`
	FreeShippingOverXValue                           float64             `bson:"free_shipping_over_x_value"`
	FreeShippingOverXIncludingTax                    bool                `bson:"free_shipping_over_x_including_tax"`
	EstimateShippingCartPageEnabled                  bool                `bson:"estimate_shipping_cart_page_enabled"`
	EstimateShippingProductPageEnabled               bool                `bson:"estimate_shipping_product_page_enabled"`
	EstimateShippingCityNameEnabled                  bool                `bson:"estimate_shipping_city_name_enabled"`
	DisplayShipmentEventsToCustomers                 bool                `bson:"display_shipment_events_to_customers"`
	DisplayShipmentEventsToStoreOwner                bool                `bson:"display_shipment_events_to_store_owner"`
	HideShippingTotal                                bool                `bson:"hide_shipping_total"`
	ShippingOriginAddressID                          int                 `bson:"shipping_origin_address_id"`
	ReturnValidOptionsIfThereAreAny                  bool                `bson:"return_valid_options_if_there_are_any"`
	BypassShippingMethodSelectionIfOnlyOne           bool                `bson:"bypass_shipping_method_selection_if_only_one"`
	UseCubeRootMethod                                bool                `bson:"use_cube_root_method"`
	ConsiderAssociatedProductsDimensions             bool                `bson:"consider_associated_products_dimensions"`
	ShipSeparatelyOneItemEach                        bool                `bson:"ship_separately_one_item_each"`
	RequestDelay                                     int                 `bson:"request_delay"`
	ShippingSorting                                  ShippingSortingEnum `bson:"shipping_sorting"`
}

// NewShippingSettings creates a new instance of ShippingSettings with default values
func NewShippingSettings() *ShippingSettings {
	return &ShippingSettings{
		ActiveShippingRateComputationMethodSystemNames: []string{},
		ActivePickupPointProviderSystemNames:           []string{},
	}
}

// ShippingSettingsRepository defines the repository interface for ShippingSettings
type ShippingSettingsRepository interface {
	CreateMany(c context.Context, items []ShippingSettings) error
	Create(c context.Context, shipping_settings *ShippingSettings) error
	Update(c context.Context, shipping_settings *ShippingSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingSettings, error)
	FetchByID(c context.Context, ID string) (ShippingSettings, error)
}

// ShippingSettingsUsecase defines the usecase interface for ShippingSettings
type ShippingSettingsUsecase interface {
	CreateMany(c context.Context, items []ShippingSettings) error
	FetchByID(c context.Context, ID string) (ShippingSettings, error)
	Create(c context.Context, shipping_settings *ShippingSettings) error
	Update(c context.Context, shipping_settings *ShippingSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShippingSettings, error)
}
