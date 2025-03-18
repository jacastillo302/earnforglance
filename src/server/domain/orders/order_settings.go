package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrderSettings = "order_settings"
)

// OrderSettings represents order settings.
type OrderSettings struct {
	ID                                                primitive.ObjectID `bson:"_id,omitempty"`
	IsReOrderAllowed                                  bool               `bson:"is_re_order_allowed"`
	MinOrderSubtotalAmount                            float64            `bson:"min_order_subtotal_amount"`
	MinOrderSubtotalAmountIncludingTax                bool               `bson:"min_order_subtotal_amount_including_tax"`
	MinOrderTotalAmount                               float64            `bson:"min_order_total_amount"`
	AutoUpdateOrderTotalsOnEditingOrder               bool               `bson:"auto_update_order_totals_on_editing_order"`
	AnonymousCheckoutAllowed                          bool               `bson:"anonymous_checkout_allowed"`
	CheckoutDisabled                                  bool               `bson:"checkout_disabled"`
	TermsOfServiceOnShoppingCartPage                  bool               `bson:"terms_of_service_on_shopping_cart_page"`
	TermsOfServiceOnOrderConfirmPage                  bool               `bson:"terms_of_service_on_order_confirm_page"`
	OnePageCheckoutEnabled                            bool               `bson:"one_page_checkout_enabled"`
	OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab bool               `bson:"one_page_checkout_display_order_totals_on_payment_info_tab"`
	DisableBillingAddressCheckoutStep                 bool               `bson:"disable_billing_address_checkout_step"`
	DisableOrderCompletedPage                         bool               `bson:"disable_order_completed_page"`
	DisplayPickupInStoreOnShippingMethodPage          bool               `bson:"display_pickup_in_store_on_shipping_method_page"`
	AttachPdfInvoiceToOrderPlacedEmail                bool               `bson:"attach_pdf_invoice_to_order_placed_email"`
	AttachPdfInvoiceToOrderPaidEmail                  bool               `bson:"attach_pdf_invoice_to_order_paid_email"`
	AttachPdfInvoiceToOrderProcessingEmail            bool               `bson:"attach_pdf_invoice_to_order_processing_email"`
	AttachPdfInvoiceToOrderCompletedEmail             bool               `bson:"attach_pdf_invoice_to_order_completed_email"`
	GeneratePdfInvoiceInCustomerLanguage              bool               `bson:"generate_pdf_invoice_in_customer_language"`
	ReturnRequestsEnabled                             bool               `bson:"return_requests_enabled"`
	ReturnRequestsAllowFiles                          bool               `bson:"return_requests_allow_files"`
	ReturnRequestsFileMaximumSize                     int                `bson:"return_requests_file_maximum_size"`
	ReturnRequestNumberMask                           string             `bson:"return_request_number_mask"`
	NumberOfDaysReturnRequestAvailable                int                `bson:"number_of_days_return_request_available"`
	ActivateGiftCardsAfterCompletingOrder             bool               `bson:"activate_gift_cards_after_completing_order"`
	DeactivateGiftCardsAfterCancellingOrder           bool               `bson:"deactivate_gift_cards_after_cancelling_order"`
	DeactivateGiftCardsAfterDeletingOrder             bool               `bson:"deactivate_gift_cards_after_deleting_order"`
	MinimumOrderPlacementInterval                     int                `bson:"minimum_order_placement_interval"`
	CompleteOrderWhenDelivered                        bool               `bson:"complete_order_when_delivered"`
	CustomOrderNumberMask                             string             `bson:"custom_order_number_mask"`
	ExportWithProducts                                bool               `bson:"export_with_products"`
	AllowAdminsToBuyCallForPriceProducts              bool               `bson:"allow_admins_to_buy_call_for_price_products"`
	ShowProductThumbnailInOrderDetailsPage            bool               `bson:"show_product_thumbnail_in_order_details_page"`
	DeleteGiftCardUsageHistory                        bool               `bson:"delete_gift_card_usage_history"`
	DisplayCustomerCurrencyOnOrders                   bool               `bson:"display_customer_currency_on_orders"`
	DisplayOrderSummary                               bool               `bson:"display_order_summary"`
	PlaceOrderWithLock                                bool               `bson:"place_order_with_lock"`
}

// OrderSettingsRepository represents the repository interface for OrderSettings
type OrderSettingsRepository interface {
	Create(c context.Context, order_settings *OrderSettings) error
	Update(c context.Context, order_settings *OrderSettings) error
	Delete(c context.Context, order_settings *OrderSettings) error
	Fetch(c context.Context) ([]OrderSettings, error)
	FetchByID(c context.Context, order_settingsID string) (OrderSettings, error)
}

// OrderSettingsUsecase represents the usecase interface for OrderSettings
type OrderSettingsUsecase interface {
	FetchByID(c context.Context, order_settingsID string) (OrderSettings, error)
	Create(c context.Context, order_settings *OrderSettings) error
	Update(c context.Context, order_settings *OrderSettings) error
	Delete(c context.Context, order_settings *OrderSettings) error
	Fetch(c context.Context) ([]OrderSettings, error)
}
