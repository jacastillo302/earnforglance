package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCustomerAddressMapping = "customer_address_mappings"
)

// CustomerAddressMapping represents a customer-address mapping class
type CustomerAddressMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID int                `bson:"customer_id"`
	AddressID  int                `bson:"address_id"`
}
