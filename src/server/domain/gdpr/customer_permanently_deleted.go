package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerPermanentlyDeleted = "customer_permanently_deletes"
)

// CustomerPermanentlyDeleted represents a customer permanently deleted (GDPR)
type CustomerPermanentlyDeleted struct {
	CustomerID primitive.ObjectID `bson:"customer_id"`
	Email      string             `bson:"email"`
}

// NewCustomerPermanentlyDeleted creates a new CustomerPermanentlyDeleted instance
func NewCustomerPermanentlyDeleted(customerID primitive.ObjectID, email string) *CustomerPermanentlyDeleted {
	return &CustomerPermanentlyDeleted{
		CustomerID: customerID,
		Email:      email,
	}
}

// CustomerPermanentlyDeletedRepository defines the repository interface for CustomerPermanentlyDeleted
type CustomerPermanentlyDeletedRepository interface {
	CreateMany(c context.Context, items []CustomerPermanentlyDeleted) error
	Create(c context.Context, customer_permanently_deleted *CustomerPermanentlyDeleted) error
	Update(c context.Context, customer_permanently_deleted *CustomerPermanentlyDeleted) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerPermanentlyDeleted, error)
	FetchByID(c context.Context, ID string) (CustomerPermanentlyDeleted, error)
}

// CustomerPermanentlyDeletedUsecase defines the usecase interface for CustomerPermanentlyDeleted
type CustomerPermanentlyDeletedUsecase interface {
	CreateMany(c context.Context, items []CustomerPermanentlyDeleted) error
	FetchByID(c context.Context, ID string) (CustomerPermanentlyDeleted, error)
	Create(c context.Context, customer_permanently_deleted *CustomerPermanentlyDeleted) error
	Update(c context.Context, customer_permanently_deleted *CustomerPermanentlyDeleted) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerPermanentlyDeleted, error)
}
