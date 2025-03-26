package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRecurringPayment = "recurring_payments"
)

// RecurringPayment represents a recurring payment.
type RecurringPayment struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`
	CycleLength                   int                `bson:"cycle_length"`
	RecurringProductCyclePeriodID int                `bson:"cycle_period_id"`
	TotalCycles                   int                `bson:"total_cycles"`
	StartDateUtc                  time.Time          `bson:"start_date_utc"`
	IsActive                      bool               `bson:"is_active"`
	LastPaymentFailed             bool               `bson:"last_payment_failed"`
	Deleted                       bool               `bson:"deleted"`
	OrderID                       int                `bson:"initial_order_id"`
	CreatedOnUtc                  time.Time          `bson:"created_on_utc"`
	CyclePeriod                   int                `bson:"cycle_period"`
}

// RecurringPaymentRepository interface
type RecurringPaymentRepository interface {
	CreateMany(c context.Context, items []RecurringPayment) error
	Create(c context.Context, recurring_payment *RecurringPayment) error
	Update(c context.Context, recurring_payment *RecurringPayment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RecurringPayment, error)
	FetchByID(c context.Context, ID string) (RecurringPayment, error)
}

// RecurringPaymentUsecase interface
type RecurringPaymentUsecase interface {
	CreateMany(c context.Context, items []RecurringPayment) error
	FetchByID(c context.Context, ID string) (RecurringPayment, error)
	Create(c context.Context, recurring_payment *RecurringPayment) error
	Update(c context.Context, recurring_payment *RecurringPayment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RecurringPayment, error)
}
