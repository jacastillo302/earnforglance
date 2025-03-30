package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerPassword = "customer_passwords"
)

// CustomerPassword represents a customer password
type CustomerPassword struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`   // MongoDB ObjectID
	CustomerID       primitive.ObjectID `bson:"customer_id"`     // Customer identifier
	Password         string             `bson:"password"`        // Password
	PasswordFormatID int                `bson:"password_format"` // Password format
	PasswordSalt     string             `bson:"password_salt"`   // Password salt
	CreatedOnUTC     time.Time          `bson:"created_on_utc"`  // Date and time of entity creation
}

type CustomerPasswordRepository interface {
	CreateMany(c context.Context, items []CustomerPassword) error
	Create(c context.Context, permission_record_customer_role_mapping *CustomerPassword) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerPassword) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerPassword, error)
	FetchByID(c context.Context, ID string) (CustomerPassword, error)
}

type CustomerPasswordUsecase interface {
	CreateMany(c context.Context, items []CustomerPassword) error
	FetchByID(c context.Context, ID string) (CustomerPassword, error)
	Create(c context.Context, permission_record_customer_role_mapping *CustomerPassword) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerPassword) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerPassword, error)
}
