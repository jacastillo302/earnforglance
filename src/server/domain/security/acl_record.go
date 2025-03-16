package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionAclRecord = "acl_records"
)

// AclRecord represents an ACL record
type AclRecord struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	EntityID       int                `bson:"entity_id"`
	EntityName     string             `bson:"entity_name"`
	CustomerRoleID int                `bson:"customer_role_id"`
}
