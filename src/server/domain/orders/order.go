package domain

import (
	"context" // added context library
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrder = "orders"
)

// Order represents an order
type Order struct {
	ID                                      primitive.ObjectID  `bson:"_id,omitempty"`
	OrderGuid                               uuid.UUID           `bson:"order_guid"`
	StoreID                                 primitive.ObjectID  `bson:"store_id"`
	CustomerID                              primitive.ObjectID  `bson:"customer_id"`
	BillingAddressID                        primitive.ObjectID  `bson:"billing_address_id"`
	ShippingAddressID                       *primitive.ObjectID `bson:"shipping_address_id,omitempty"`
	PickupAddressID                         *primitive.ObjectID `bson:"pickup_address_id,omitempty"`
	PickupInStore                           bool                `bson:"pickup_in_store"`
	OrderStatusID                           int                 `bson:"order_status_id"`
	ShippingStatusID                        int                 `bson:"shipping_status_id"`
	PaymentStatusID                         int                 `bson:"payment_status_id"`
	PaymentMethodSystemName                 string              `bson:"payment_method_system_name"`
	CustomerCurrencyCode                    string              `bson:"customer_currency_code"`
	CurrencyRate                            float64             `bson:"currency_rate"`
	CustomerTaxDisplayTypeID                int                 `bson:"customer_tax_display_type_id"`
	VatNumber                               string              `bson:"vat_number"`
	OrderSubtotalInclTax                    float64             `bson:"order_subtotal_incl_tax"`
	OrderSubtotalExclTax                    float64             `bson:"order_subtotal_excl_tax"`
	OrderSubTotalDiscountInclTax            float64             `bson:"order_subtotal_discount_incl_tax"`
	OrderSubTotalDiscountExclTax            float64             `bson:"order_subtotal_discount_excl_tax"`
	OrderShippingInclTax                    float64             `bson:"order_shipping_incl_tax"`
	OrderShippingExclTax                    float64             `bson:"order_shipping_excl_tax"`
	PaymentMethodAdditionalFeeInclTax       float64             `bson:"payment_method_additional_fee_incl_tax"`
	PaymentMethodAdditionalFeeExclTax       float64             `bson:"payment_method_additional_fee_excl_tax"`
	TaxRates                                string              `bson:"tax_rates"`
	OrderTax                                float64             `bson:"order_tax"`
	OrderDiscount                           float64             `bson:"order_discount"`
	OrderTotal                              float64             `bson:"order_total"`
	RefundedAmount                          float64             `bson:"refunded_amount"`
	RewardPointsHistoryEntryID              *primitive.ObjectID `bson:"reward_points_history_entry_id,omitempty"`
	CheckoutAttributeDescription            string              `bson:"checkout_attribute_description"`
	CheckoutAttributesXml                   string              `bson:"checkout_attributes_xml"`
	CustomerLanguageID                      primitive.ObjectID  `bson:"customer_language_id"`
	AffiliateID                             primitive.ObjectID  `bson:"affiliate_id"`
	CustomerIp                              string              `bson:"customer_ip"`
	AllowStoringCreditCardNumber            bool                `bson:"allow_storing_credit_card_number"`
	CardType                                string              `bson:"card_type"`
	CardName                                string              `bson:"card_name"`
	CardNumber                              string              `bson:"card_number"`
	MaskedCreditCardNumber                  string              `bson:"masked_credit_card_number"`
	CardCvv2                                string              `bson:"card_cvv2"`
	CardExpirationMonth                     string              `bson:"card_expiration_month"`
	CardExpirationYear                      string              `bson:"card_expiration_year"`
	AuthorizationTransactionID              string              `bson:"authorization_transaction_id"`
	AuthorizationTransactionCode            string              `bson:"authorization_transaction_code"`
	AuthorizationTransactionResult          string              `bson:"authorization_transaction_result"`
	CaptureTransactionID                    string              `bson:"capture_transaction_id"`
	CaptureTransactionResult                string              `bson:"capture_transaction_result"`
	SubscriptionTransactionID               string              `bson:"subscription_transaction_id"`
	PaidDateUtc                             *time.Time          `bson:"paid_date_utc,omitempty"`
	ShippingMethod                          string              `bson:"shipping_method"`
	ShippingRateComputationMethodSystemName string              `bson:"shipping_rate_computation_method_system_name"`
	CustomValuesXml                         string              `bson:"custom_values_xml"`
	Deleted                                 bool                `bson:"deleted"`
	CreatedOnUtc                            time.Time           `bson:"created_on_utc"`
	CustomOrderNumber                       string              `bson:"custom_order_number"`
	RedeemedRewardPointsEntryID             *primitive.ObjectID `bson:"redeemed_reward_points_entry_id,omitempty"`
	OrderStatus                             OrderStatus         `bson:"order_status"`
	PaymentStatus                           int                 `bson:"payment_status"`
	ShippingStatus                          int                 `bson:"shipping_status"`
	CustomerTaxDisplayType                  int                 `bson:"customer_tax_display_type"`
}

// OrderRepository represents the order repository interface
type OrderRepository interface {
	CreateMany(c context.Context, items []Order) error
	Create(c context.Context, order *Order) error
	Update(c context.Context, order *Order) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Order, error)
	FetchByID(c context.Context, ID string) (Order, error)
}

// OrderUsecase represents the order usecase interface
type OrderUsecase interface {
	CreateMany(c context.Context, items []Order) error
	FetchByID(c context.Context, ID string) (Order, error)
	Create(c context.Context, order *Order) error
	Update(c context.Context, order *Order) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Order, error)
}
