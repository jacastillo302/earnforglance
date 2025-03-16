package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRecurringPayment = "recurring_payments"
)

// RecurringPayment represents a recurring payment
type RecurringPayment struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	CycleLength       int                `bson:"cycle_length"`
	CyclePeriodID     int                `bson:"cycle_period_id"`
	TotalCycles       int                `bson:"total_cycles"`
	StartDateUtc      time.Time          `bson:"start_date_utc"`
	IsActive          bool               `bson:"is_active"`
	LastPaymentFailed bool               `bson:"last_payment_failed"`
	Deleted           bool               `bson:"deleted"`
	InitialOrderID    int                `bson:"initial_order_id"`
	CreatedOnUtc      time.Time          `bson:"created_on_utc"`
	CyclePeriod       int                `bson:"cycle_period"`
}
