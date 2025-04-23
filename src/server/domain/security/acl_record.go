package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionAclRecord = "acl_records"
)

// AclRecord represents an ACL record.
type AclRecord struct {
	EntityID       bson.ObjectID `bson:"_id,omitempty"`
	EntityName     string        `bson:"entity_name"`
	CustomerRoleID bson.ObjectID `bson:"customer_role_id"`
	IsRead         bool          `bson:"is_read"`
	IsDelete       bool          `bson:"is_delete"`
	IsUpdate       bool          `bson:"is_update"`
	IsCreate       bool          `bson:"is_create"`
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
	CreateMany(c context.Context, items []AclRecord) error
	FetchByID(c context.Context, ID string) (AclRecord, error)
	Create(c context.Context, acl_record *AclRecord) error
	Update(c context.Context, acl_record *AclRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AclRecord, error)
}
