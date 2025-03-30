package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerRole = "customer_roles" // MongoDB collection name for customer roles
)

// CustomerRole represents a customer role
type CustomerRole struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`               // MongoDB ObjectID
	Name                    string             `bson:"name"`                        // Customer role name
	FreeShipping            bool               `bson:"free_shipping"`               // Indicates if the role is marked as free shipping
	TaxExempt               bool               `bson:"tax_exempt"`                  // Indicates if the role is marked as tax exempt
	Active                  bool               `bson:"active"`                      // Indicates if the role is active
	IsSystemRole            bool               `bson:"is_system_role"`              // Indicates if the role is a system role
	SystemName              string             `bson:"system_name"`                 // System name of the role
	EnablePasswordLifetime  bool               `bson:"enable_password_lifetime"`    // Indicates if customers must change passwords after a specified time
	OverrideTaxDisplayType  bool               `bson:"override_tax_display_type"`   // Indicates if the role has a custom tax display type
	DefaultTaxDisplayTypeID int                `bson:"default_tax_display_type_id"` // Identifier of the default tax display type
	PurchasedWithProductId  int                `bson:"purchased_with_product_id"`   // Product ID required for this role
}

type CustomerRoleRepository interface {
	CreateMany(c context.Context, items []CustomerRole) error
	Create(c context.Context, permission_record_customer_role_mapping *CustomerRole) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerRole) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerRole, error)
	FetchByID(c context.Context, ID string) (CustomerRole, error)
}

type CustomerRoleUsecase interface {
	CreateMany(c context.Context, items []CustomerRole) error
	FetchByID(c context.Context, ID string) (CustomerRole, error)
	Create(c context.Context, permission_record_customer_role_mapping *CustomerRole) error
	Update(c context.Context, permission_record_customer_role_mapping *CustomerRole) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerRole, error)
}
