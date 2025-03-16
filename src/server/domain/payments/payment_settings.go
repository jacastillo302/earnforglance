package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPaymentSettings = "payment_settings"
)

// PaymentSettings represents payment settings
type PaymentSettings struct {
	ID                                              primitive.ObjectID `bson:"_id,omitempty"`
	ActivePaymentMethodSystemNames                  []string           `bson:"active_payment_method_system_names"`
	AllowRePostingPayments                          bool               `bson:"allow_re_posting_payments"`
	BypassPaymentMethodSelectionIfOnlyOne           bool               `bson:"bypass_payment_method_selection_if_only_one"`
	ShowPaymentMethodDescriptions                   bool               `bson:"show_payment_method_descriptions"`
	SkipPaymentInfoStepForRedirectionPaymentMethods bool               `bson:"skip_payment_info_step_for_redirection_payment_methods"`
	CancelRecurringPaymentsAfterFailedPayment       bool               `bson:"cancel_recurring_payments_after_failed_payment"`
	RegenerateOrderGuidInterval                     int                `bson:"regenerate_order_guid_interval"`
}

// NewPaymentSettings creates a new instance of PaymentSettings with default values
func NewPaymentSettings() *PaymentSettings {
	return &PaymentSettings{
		ActivePaymentMethodSystemNames: []string{},
	}
}
