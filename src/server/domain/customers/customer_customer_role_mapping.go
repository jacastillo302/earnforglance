package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCustomerCustomerRoleMapping = "customer_customer_role_mappings"
)

// CustomerCustomerRoleMapping represents a customer-customer role mapping class
type CustomerCustomerRoleMapping struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID     int                `bson:"customer_id"`
	CustomerRoleID int                `bson:"customer_role_id"`
}
