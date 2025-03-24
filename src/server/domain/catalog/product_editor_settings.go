package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductEditorSettings = "product_editor_settings"
)

// ProductEditorSettings represents product editor settings
type ProductEditorSettings struct {
	ID                                           primitive.ObjectID `bson:"_id,omitempty"`
	ProductType                                  bool               `bson:"product_type"`
	VisibleIndividually                          bool               `bson:"visible_individually"`
	ProductTemplate                              bool               `bson:"product_template"`
	AdminComment                                 bool               `bson:"admin_comment"`
	Vendor                                       bool               `bson:"vendor"`
	Stores                                       bool               `bson:"stores"`
	ACL                                          bool               `bson:"acl"`
	ShowOnHomepage                               bool               `bson:"show_on_homepage"`
	AllowCustomerReviews                         bool               `bson:"allow_customer_reviews"`
	ProductTags                                  bool               `bson:"product_tags"`
	ManufacturerPartNumber                       bool               `bson:"manufacturer_part_number"`
	GTIN                                         bool               `bson:"gtin"`
	ProductCost                                  bool               `bson:"product_cost"`
	TierPrices                                   bool               `bson:"tier_prices"`
	Discounts                                    bool               `bson:"discounts"`
	DisableBuyButton                             bool               `bson:"disable_buy_button"`
	DisableWishlistButton                        bool               `bson:"disable_wishlist_button"`
	AvailableForPreOrder                         bool               `bson:"available_for_pre_order"`
	CallForPrice                                 bool               `bson:"call_for_price"`
	OldPrice                                     bool               `bson:"old_price"`
	CustomerEntersPrice                          bool               `bson:"customer_enters_price"`
	PAngV                                        bool               `bson:"pangv"`
	RequireOtherProductsAddedToCart              bool               `bson:"require_other_products_added_to_cart"`
	IsGiftCard                                   bool               `bson:"is_gift_card"`
	DownloadableProduct                          bool               `bson:"downloadable_product"`
	RecurringProduct                             bool               `bson:"recurring_product"`
	IsRental                                     bool               `bson:"is_rental"`
	FreeShipping                                 bool               `bson:"free_shipping"`
	ShipSeparately                               bool               `bson:"ship_separately"`
	AdditionalShippingCharge                     bool               `bson:"additional_shipping_charge"`
	DeliveryDate                                 bool               `bson:"delivery_date"`
	ProductAvailabilityRange                     bool               `bson:"product_availability_range"`
	UseMultipleWarehouses                        bool               `bson:"use_multiple_warehouses"`
	Warehouse                                    bool               `bson:"warehouse"`
	DisplayStockAvailability                     bool               `bson:"display_stock_availability"`
	MinimumStockQuantity                         bool               `bson:"minimum_stock_quantity"`
	LowStockActivity                             bool               `bson:"low_stock_activity"`
	NotifyAdminForQuantityBelow                  bool               `bson:"notify_admin_for_quantity_below"`
	Backorders                                   bool               `bson:"backorders"`
	AllowBackInStockSubscriptions                bool               `bson:"allow_back_in_stock_subscriptions"`
	MinimumCartQuantity                          bool               `bson:"minimum_cart_quantity"`
	MaximumCartQuantity                          bool               `bson:"maximum_cart_quantity"`
	AllowedQuantities                            bool               `bson:"allowed_quantities"`
	AllowAddingOnlyExistingAttributeCombinations bool               `bson:"allow_adding_only_existing_attribute_combinations"`
	NotReturnable                                bool               `bson:"not_returnable"`
	Weight                                       bool               `bson:"weight"`
	Dimensions                                   bool               `bson:"dimensions"`
	AvailableStartDate                           bool               `bson:"available_start_date"`
	AvailableEndDate                             bool               `bson:"available_end_date"`
	MarkAsNew                                    bool               `bson:"mark_as_new"`
	Published                                    bool               `bson:"published"`
	RelatedProducts                              bool               `bson:"related_products"`
	CrossSellsProducts                           bool               `bson:"cross_sells_products"`
	Seo                                          bool               `bson:"seo"`
	PurchasedWithOrders                          bool               `bson:"purchased_with_orders"`
	ProductAttributes                            bool               `bson:"product_attributes"`
	SpecificationAttributes                      bool               `bson:"specification_attributes"`
	Manufacturers                                bool               `bson:"manufacturers"`
	StockQuantityChange                          bool               `bson:"stock_quantity_history"`
	AgeVerification                              bool               `bson:"age_verification"`
}

type ProductEditorSettingsRepository interface {
	CreateMany(c context.Context, items []ProductEditorSettings) error
	Create(c context.Context, product_editor_settings *ProductEditorSettings) error
	Update(c context.Context, product_editor_settings *ProductEditorSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductEditorSettings, error)
	FetchByID(c context.Context, ID string) (ProductEditorSettings, error)
}

type ProductEditorSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (ProductEditorSettings, error)
	Create(c context.Context, product_editor_settings *ProductEditorSettings) error
	Update(c context.Context, product_editor_settings *ProductEditorSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProductEditorSettings, error)
}
