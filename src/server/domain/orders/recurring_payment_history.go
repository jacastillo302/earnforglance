package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRecurringPaymentHistory = "recurring_payment_histories"
)

// RecurringPaymentHistory represents a recurring payment history
type RecurringPaymentHistory struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	RecurringPaymentID int                `bson:"recurring_payment_id"`
	OrderID            int                `bson:"order_id"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
}
