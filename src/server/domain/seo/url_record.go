package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionUrlRecord = "url_records"
)

// UrlRecord represents an URL record
type UrlRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EntityID   int                `bson:"entity_id"`
	EntityName string             `bson:"entity_name"`
	Slug       string             `bson:"slug"`
	IsActive   bool               `bson:"is_active"`
	LanguageID int                `bson:"language_id"`
}
