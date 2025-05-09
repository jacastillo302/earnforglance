package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCustomerAddressMapping = "customer_address_mappings"
)

// CustomerAddressMapping represents a customer-address mapping class
type CustomerAddressMapping struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	CustomerID bson.ObjectID `bson:"customer_id"`
	AddressID  bson.ObjectID `bson:"address_id"`
}

type CustomerAddressMappingRepository interface {
	CreateMany(c context.Context, items []CustomerAddressMapping) error
	Create(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Update(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAddressMapping, error)
	FetchByID(c context.Context, ID string) (CustomerAddressMapping, error)
}

type CustomerAddressMappingUsecase interface {
	CreateMany(c context.Context, items []CustomerAddressMapping) error
	FetchByID(c context.Context, ID string) (CustomerAddressMapping, error)
	Create(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Update(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAddressMapping, error)
}
