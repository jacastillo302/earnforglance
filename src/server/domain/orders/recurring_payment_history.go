package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionRecurringPaymentHistory = "recurring_payment_histories"
)

// RecurringPaymentHistory represents a recurring payment history
type RecurringPaymentHistory struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	RecurringPaymentID bson.ObjectID `bson:"recurring_payment_id"`
	OrderID            bson.ObjectID `bson:"order_id"`
	CreatedOnUtc       time.Time     `bson:"created_on_utc"`
}

// RecurringPaymentHistoryRepository represents the repository interface for RecurringPaymentHistory
type RecurringPaymentHistoryRepository interface {
	CreateMany(c context.Context, items []RecurringPaymentHistory) error
	Create(c context.Context, recurring_payment_history *RecurringPaymentHistory) error
	Update(c context.Context, recurring_payment_history *RecurringPaymentHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RecurringPaymentHistory, error)
	FetchByID(c context.Context, ID string) (RecurringPaymentHistory, error)
}

// RecurringPaymentHistoryUsecase represents the use case interface for RecurringPaymentHistory
type RecurringPaymentHistoryUsecase interface {
	CreateMany(c context.Context, items []RecurringPaymentHistory) error
	FetchByID(c context.Context, ID string) (RecurringPaymentHistory, error)
	Create(c context.Context, recurring_payment_history *RecurringPaymentHistory) error
	Update(c context.Context, recurring_payment_history *RecurringPaymentHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RecurringPaymentHistory, error)
}
