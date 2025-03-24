package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAclRecord = "acl_records"
)

// AclRecord represents an ACL record.
type AclRecord struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	EntityID       primitive.ObjectID `bson:"entity_id"`
	EntityName     string             `bson:"entity_name"`
	CustomerRoleID primitive.ObjectID `bson:"customer_role_id"`
}

// AclRecordRepository defines the repository interface for AclRecord
type AclRecordRepository interface {
	CreateMany(c context.Context, items []AclRecord) error
	Create(c context.Context, acl_record *AclRecord) error
	Update(c context.Context, acl_record *AclRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AclRecord, error)
	FetchByID(c context.Context, ID string) (AclRecord, error)
}

// AclRecordUsecase defines the use case interface for AclRecord
type AclRecordUsecase interface {
	FetchByID(c context.Context, ID string) (AclRecord, error)
	Create(c context.Context, acl_record *AclRecord) error
	Update(c context.Context, acl_record *AclRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AclRecord, error)
}
