package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionUrlRecord = "url_records"
)

// UrlRecord represents an URL record.
type UrlRecord struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	PermissionRecordID bson.ObjectID `bson:"entity_id"`
	Slug               string        `bson:"slug"`
	IsActive           bool          `bson:"is_active"`
}

// UrlRecordRepository defines the repository interface for UrlRecord
type UrlRecordRepository interface {
	CreateMany(c context.Context, items []UrlRecord) error
	Create(c context.Context, url_record *UrlRecord) error
	Update(c context.Context, url_record *UrlRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]UrlRecord, error)
	FetchByID(c context.Context, ID string) (UrlRecord, error)
}

// UrlRecordUsecase defines the use case interface for UrlRecord
type UrlRecordUsecase interface {
	CreateMany(c context.Context, items []UrlRecord) error
	FetchByID(c context.Context, ID string) (UrlRecord, error)
	Create(c context.Context, url_record *UrlRecord) error
	Update(c context.Context, url_record *UrlRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]UrlRecord, error)
}
