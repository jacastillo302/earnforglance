package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPermissionRecordCustomerRoleMapping = "permission_record_customer_role_mappings"
)

// PermissionRecordCustomerRoleMapping represents a permission record-customer role mapping class
type PermissionRecordCustomerRoleMapping struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	PermissionRecordID bson.ObjectID `bson:"permission_record_id"`
	CustomerRoleID     bson.ObjectID `bson:"customer_role_id"`
	IsRead             bool          `bson:"is_read"`
	IsDelete           bool          `bson:"is_delete"`
	IsUpdate           bool          `bson:"is_update"`
	IsCreate           bool          `bson:"is_create"`
}

type PermissionRecordCustomerRoleMappingRepository interface {
	CreateMany(c context.Context, items []PermissionRecordCustomerRoleMapping) error
	Create(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Update(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermissionRecordCustomerRoleMapping, error)
	FetchByID(c context.Context, ID string) (PermissionRecordCustomerRoleMapping, error)
}

type PermissionRecordCustomerRoleMappingUsecase interface {
	CreateMany(c context.Context, items []PermissionRecordCustomerRoleMapping) error
	FetchByID(c context.Context, ID string) (PermissionRecordCustomerRoleMapping, error)
	Create(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Update(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermissionRecordCustomerRoleMapping, error)
}
