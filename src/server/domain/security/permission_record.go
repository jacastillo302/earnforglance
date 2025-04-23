package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPermissionRecord = "permission_records"
)

// PermissionRecord represents a permission record.
type PermissionRecord struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	SystemName string        `bson:"system_name"`
	Category   string        `bson:"category"`
}

// PermissionRecordRepository defines the repository interface for PermissionRecord
type PermissionRecordRepository interface {
	CreateMany(c context.Context, items []PermissionRecord) error
	Create(c context.Context, permission_record *PermissionRecord) error
	Update(c context.Context, permission_record *PermissionRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermissionRecord, error)
	FetchByID(c context.Context, ID string) (PermissionRecord, error)
}

// PermissionRecordUsecase defines the usecase interface for PermissionRecord
type PermissionRecordUsecase interface {
	CreateMany(c context.Context, items []PermissionRecord) error
	FetchByID(c context.Context, ID string) (PermissionRecord, error)
	Create(c context.Context, permission_record *PermissionRecord) error
	Update(c context.Context, permission_record *PermissionRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermissionRecord, error)
}
