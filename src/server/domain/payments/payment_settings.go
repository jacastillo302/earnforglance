package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPaymentSettings = "payment_settings"
)

// PaymentSettings represents payment settings.
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

type PaymentSettingsRepository interface {
	CreateMany(c context.Context, items []PaymentSettings) error
	Create(c context.Context, payment_settings *PaymentSettings) error
	Update(c context.Context, payment_settings *PaymentSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PaymentSettings, error)
	FetchByID(c context.Context, ID string) (PaymentSettings, error)
}

type PaymentSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (PaymentSettings, error)
	Create(c context.Context, payment_settings *PaymentSettings) error
	Update(c context.Context, payment_settings *PaymentSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PaymentSettings, error)
}
