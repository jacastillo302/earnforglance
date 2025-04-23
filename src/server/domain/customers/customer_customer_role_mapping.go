package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCustomerCustomerRoleMapping = "customer_customer_role_mappings"
)

// CustomerCustomerRoleMapping represents a customer-customer role mapping class
type CustomerCustomerRoleMapping struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	CustomerID     bson.ObjectID `bson:"customer_id,omitempty"`
	CustomerRoleID bson.ObjectID `bson:"customer_role_id,omitempty"`
}

type CustomerCustomerRoleMappingRepository interface {
	CreateMany(c context.Context, items []CustomerCustomerRoleMapping) error
	Create(c context.Context, customer_customer_role_mapping *CustomerCustomerRoleMapping) error
	Update(c context.Context, customer_customer_role_mapping *CustomerCustomerRoleMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerCustomerRoleMapping, error)
	FetchByID(c context.Context, ID string) (CustomerCustomerRoleMapping, error)
}

type CustomerCustomerRoleMappingUsecase interface {
	CreateMany(c context.Context, items []CustomerCustomerRoleMapping) error
	FetchByID(c context.Context, ID string) (CustomerCustomerRoleMapping, error)
	Create(c context.Context, customer_customer_role_mapping *CustomerCustomerRoleMapping) error
	Update(c context.Context, customer_customer_role_mapping *CustomerCustomerRoleMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerCustomerRoleMapping, error)
}
