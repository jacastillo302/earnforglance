package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerAddressMapping = "customer_address_mappings"
)

// CustomerAddressMapping represents a customer-address mapping class
type CustomerAddressMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id"`
	AddressID  primitive.ObjectID `bson:"address_id"`
}

type CustomerAddressMappingRepository interface {
	Create(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Update(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAddressMapping, error)
	FetchByID(c context.Context, ID string) (CustomerAddressMapping, error)
}

type CustomerAddressMappingUsecase interface {
	FetchByID(c context.Context, ID string) (CustomerAddressMapping, error)
	Create(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Update(c context.Context, customer_address_mapping *CustomerAddressMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAddressMapping, error)
}
