package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProduct = "products"
)

// Product represents a product
type Product struct {
	ID                                           primitive.ObjectID          `bson:"_id,omitempty"`
	ProductTypeID                                int                         `bson:"product_type_id"`
	ParentGroupedProductID                       primitive.ObjectID          `bson:"parent_grouped_product_id"`
	VisibleIndividually                          bool                        `bson:"visible_individually"`
	Name                                         string                      `bson:"name"`
	ShortDescription                             string                      `bson:"short_description"`
	FullDescription                              string                      `bson:"full_description"`
	AdminComment                                 string                      `bson:"admin_comment"`
	ProductTemplateID                            primitive.ObjectID          `bson:"product_template_id"`
	VendorID                                     primitive.ObjectID          `bson:"vendor_id"`
	ShowOnHomepage                               bool                        `bson:"show_on_homepage"`
	MetaKeywords                                 string                      `bson:"meta_keywords"`
	MetaDescription                              string                      `bson:"meta_description"`
	MetaTitle                                    string                      `bson:"meta_title"`
	AllowCustomerReviews                         bool                        `bson:"allow_customer_reviews"`
	ApprovedRatingSum                            int                         `bson:"approved_rating_sum"`
	NotApprovedRatingSum                         int                         `bson:"not_approved_rating_sum"`
	ApprovedTotalReviews                         int                         `bson:"approved_total_reviews"`
	NotApprovedTotalReviews                      int                         `bson:"not_approved_total_reviews"`
	SubjectToAcl                                 bool                        `bson:"subject_to_acl"`
	LimitedToStores                              bool                        `bson:"limited_to_stores"`
	Sku                                          string                      `bson:"sku"`
	ManufacturerPartNumber                       string                      `bson:"manufacturer_part_number"`
	Gtin                                         string                      `bson:"gtin"`
	IsGiftCard                                   bool                        `bson:"is_gift_card"`
	GiftCardTypeID                               int                         `bson:"gift_card_type_id"`
	OverriddenGiftCardAmount                     *float64                    `bson:"overridden_gift_card_amount,omitempty"`
	RequireOtherProducts                         bool                        `bson:"require_other_products"`
	RequiredProductIds                           string                      `bson:"required_product_ids"`
	AutomaticallyAddRequiredProducts             bool                        `bson:"automatically_add_required_products"`
	IsDownload                                   bool                        `bson:"is_download"`
	DownloadID                                   int                         `bson:"download_id"`
	UnlimitedDownloads                           bool                        `bson:"unlimited_downloads"`
	MaxNumberOfDownloads                         int                         `bson:"max_number_of_downloads"`
	DownloadExpirationDays                       *int                        `bson:"download_expiration_days,omitempty"`
	DownloadActivationTypeID                     int                         `bson:"download_activation_type_id"`
	HasSampleDownload                            bool                        `bson:"has_sample_download"`
	SampleDownloadID                             int                         `bson:"sample_download_id"`
	HasUserAgreement                             bool                        `bson:"has_user_agreement"`
	UserAgreementText                            string                      `bson:"user_agreement_text"`
	IsRecurring                                  bool                        `bson:"is_recurring"`
	RecurringCycleLength                         int                         `bson:"recurring_cycle_length"`
	RecurringCyclePeriodID                       int                         `bson:"recurring_cycle_period_id"`
	RecurringTotalCycles                         int                         `bson:"recurring_total_cycles"`
	IsRental                                     bool                        `bson:"is_rental"`
	RentalPriceLength                            int                         `bson:"rental_price_length"`
	RentalPricePeriodID                          int                         `bson:"rental_price_period_id"`
	IsShipEnabled                                bool                        `bson:"is_ship_enabled"`
	IsFreeShipping                               bool                        `bson:"is_free_shipping"`
	ShipSeparately                               bool                        `bson:"ship_separately"`
	AdditionalShippingCharge                     float64                     `bson:"additional_shipping_charge"`
	DeliveryDateID                               primitive.ObjectID          `bson:"delivery_date_id"`
	IsTaxExempt                                  bool                        `bson:"is_tax_exempt"`
	TaxCategoryID                                primitive.ObjectID          `bson:"tax_category_id"`
	ManageInventoryMethodID                      int                         `bson:"manage_inventory_method_id"`
	ProductAvailabilityRangeID                   primitive.ObjectID          `bson:"product_availability_range_id"`
	UseMultipleWarehouses                        bool                        `bson:"use_multiple_warehouses"`
	WarehouseID                                  primitive.ObjectID          `bson:"warehouse_id"`
	StockQuantity                                int                         `bson:"stock_quantity"`
	DisplayStockAvailability                     bool                        `bson:"display_stock_availability"`
	DisplayStockQuantity                         bool                        `bson:"display_stock_quantity"`
	MinStockQuantity                             int                         `bson:"min_stock_quantity"`
	LowStockActivityID                           int                         `bson:"low_stock_activity_id"`
	NotifyAdminForQuantityBelow                  int                         `bson:"notify_admin_for_quantity_below"`
	BackorderModeID                              int                         `bson:"backorder_mode_id"`
	AllowBackInStockSubscriptions                bool                        `bson:"allow_back_in_stock_subscriptions"`
	OrderMinimumQuantity                         int                         `bson:"order_minimum_quantity"`
	OrderMaximumQuantity                         int                         `bson:"order_maximum_quantity"`
	AllowedQuantities                            string                      `bson:"allowed_quantities"`
	AllowAddingOnlyExistingAttributeCombinations bool                        `bson:"allow_adding_only_existing_attribute_combinations"`
	DisplayAttributeCombinationImagesOnly        bool                        `bson:"display_attribute_combination_images_only"`
	NotReturnable                                bool                        `bson:"not_returnable"`
	DisableBuyButton                             bool                        `bson:"disable_buy_button"`
	DisableWishlistButton                        bool                        `bson:"disable_wishlist_button"`
	AvailableForPreOrder                         bool                        `bson:"available_for_pre_order"`
	PreOrderAvailabilityStartDateTimeUtc         *time.Time                  `bson:"pre_order_availability_start_date_time_utc,omitempty"`
	CallForPrice                                 bool                        `bson:"call_for_price"`
	Price                                        float64                     `bson:"price"`
	OldPrice                                     float64                     `bson:"old_price"`
	ProductCost                                  float64                     `bson:"product_cost"`
	CustomerEntersPrice                          bool                        `bson:"customer_enters_price"`
	MinimumCustomerEnteredPrice                  float64                     `bson:"minimum_customer_entered_price"`
	MaximumCustomerEnteredPrice                  float64                     `bson:"maximum_customer_entered_price"`
	BasepriceEnabled                             bool                        `bson:"baseprice_enabled"`
	BasepriceAmount                              float64                     `bson:"baseprice_amount"`
	BasepriceUnitID                              int                         `bson:"baseprice_unit_id"`
	BasepriceBaseAmount                          float64                     `bson:"baseprice_base_amount"`
	BasepriceBaseUnitID                          int                         `bson:"baseprice_base_unit_id"`
	MarkAsNew                                    bool                        `bson:"mark_as_new"`
	MarkAsNewStartDateTimeUtc                    *time.Time                  `bson:"mark_as_new_start_date_time_utc,omitempty"`
	MarkAsNewEndDateTimeUtc                      *time.Time                  `bson:"mark_as_new_end_date_time_utc,omitempty"`
	Weight                                       float64                     `bson:"weight"`
	Length                                       float64                     `bson:"length"`
	Width                                        float64                     `bson:"width"`
	Height                                       float64                     `bson:"height"`
	AvailableStartDateTimeUtc                    *time.Time                  `bson:"available_start_date_time_utc,omitempty"`
	AvailableEndDateTimeUtc                      *time.Time                  `bson:"available_end_date_time_utc,omitempty"`
	DisplayOrder                                 int                         `bson:"display_order"`
	Published                                    bool                        `bson:"published"`
	Deleted                                      bool                        `bson:"deleted"`
	CreatedOnUtc                                 time.Time                   `bson:"created_on_utc"`
	UpdatedOnUtc                                 time.Time                   `bson:"updated_on_utc"`
	AgeVerification                              bool                        `bson:"age_verification"`
	MinimumAgeToPurchase                         int                         `bson:"minimum_age_to_purchase"`
	ProductType                                  ProductType                 `bson:"product_type"`
	BackorderMode                                BackorderMode               `bson:"backorder_mode"`
	DownloadActivationType                       DownloadActivationType      `bson:"download_activation_type"`
	GiftCardType                                 GiftCardType                `bson:"gift_card_type"`
	LowStockActivity                             LowStockActivity            `bson:"low_stock_activity"`
	ManageInventoryMethod                        ManageInventoryMethod       `bson:"manage_inventory_method"`
	RecurringCyclePeriod                         RecurringProductCyclePeriod `bson:"recurring_cycle_period"`
	RentalPricePeriod                            RentalPricePeriod           `bson:"rental_price_period"`
}

type ProductRepository interface {
	Create(c context.Context, product *Product) error
	Update(c context.Context, product *Product) error
	Delete(c context.Context, product *Product) error
	Fetch(c context.Context) ([]Product, error)
	FetchByID(c context.Context, productID string) (Product, error)
}

type ProductUsecase interface {
	FetchByID(c context.Context, productID string) (Product, error)
	Create(c context.Context, product *Product) error
	Update(c context.Context, product *Product) error
	Delete(c context.Context, product *Product) error
	Fetch(c context.Context) ([]Product, error)
}
