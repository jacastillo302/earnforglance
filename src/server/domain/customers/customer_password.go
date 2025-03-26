package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerPassword = "customer_passwords"
)

// CustomerPassword represents a customer password
type CustomerPassword struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`   // MongoDB ObjectID
	CustomerID       int                `bson:"customer_id"`     // Customer identifier
	Password         string             `bson:"password"`        // Password
	PasswordFormatID int                `bson:"password_format"` // Password format
	PasswordSalt     string             `bson:"password_salt"`   // Password salt
	CreatedOnUTC     time.Time          `bson:"created_on_utc"`  // Date and time of entity creation
}
