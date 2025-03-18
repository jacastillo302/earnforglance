package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPermissionRecordCustomerRoleMapping = "permission_record_customer_role_mappings"
)

// PermissionRecordCustomerRoleMapping represents a permission record-customer role mapping class
type PermissionRecordCustomerRoleMapping struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	PermissionRecordID primitive.ObjectID `bson:"permission_record_id"`
	CustomerRoleID     primitive.ObjectID `bson:"customer_role_id"`
}

type PermissionRecordCustomerRoleMappingRepository interface {
	Create(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Update(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Delete(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Fetch(c context.Context) ([]PermissionRecordCustomerRoleMapping, error)
	FetchByID(c context.Context, permission_record_customer_role_mappingID string) (PermissionRecordCustomerRoleMapping, error)
}

type PermissionRecordCustomerRoleMappingUsecase interface {
	FetchByID(c context.Context, permission_record_customer_role_mappingID string) (PermissionRecordCustomerRoleMapping, error)
	Create(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Update(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Delete(c context.Context, permission_record_customer_role_mapping *PermissionRecordCustomerRoleMapping) error
	Fetch(c context.Context) ([]PermissionRecordCustomerRoleMapping, error)
}
