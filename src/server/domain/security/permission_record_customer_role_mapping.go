package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPermissionRecordCustomerRoleMapping = "permission_record_customer_role_mappings"
)

// PermissionRecordCustomerRoleMapping represents a permission record-customer role mapping class
type PermissionRecordCustomerRoleMapping struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	PermissionRecordID int                `bson:"permission_record_id"`
	CustomerRoleID     int                `bson:"customer_role_id"`
}
