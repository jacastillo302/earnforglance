package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRecurringPaymentHistory = "recurring_payment_histories"
)

// RecurringPaymentHistory represents a recurring payment history
type RecurringPaymentHistory struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	RecurringPaymentID primitive.ObjectID `bson:"recurring_payment_id"`
	OrderID            primitive.ObjectID `bson:"order_id"`
	CreatedOnUtc       time.Time          `bson:"created_on_utc"`
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
