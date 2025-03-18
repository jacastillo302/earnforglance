package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUrlRecord = "url_records"
)

// UrlRecord represents an URL record.
type UrlRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EntityID   primitive.ObjectID `bson:"entity_id"`
	EntityName string             `bson:"entity_name"`
	Slug       string             `bson:"slug"`
	IsActive   bool               `bson:"is_active"`
	LanguageID primitive.ObjectID `bson:"language_id"`
}

// UrlRecordRepository defines the repository interface for UrlRecord
type UrlRecordRepository interface {
	Create(c context.Context, url_record *UrlRecord) error
	Update(c context.Context, url_record *UrlRecord) error
	Delete(c context.Context, url_record *UrlRecord) error
	Fetch(c context.Context) ([]UrlRecord, error)
	FetchByID(c context.Context, url_recordID string) (UrlRecord, error)
}

// UrlRecordUsecase defines the use case interface for UrlRecord
type UrlRecordUsecase interface {
	FetchByID(c context.Context, url_recordID string) (UrlRecord, error)
	Create(c context.Context, url_record *UrlRecord) error
	Update(c context.Context, url_record *UrlRecord) error
	Delete(c context.Context, url_record *UrlRecord) error
	Fetch(c context.Context) ([]UrlRecord, error)
}
