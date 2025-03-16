package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSearchTerm = "search_terms"
)

// SearchTerm represents a search term record (for statistics)
type SearchTerm struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Keyword string             `bson:"keyword"`
	StoreID int                `bson:"store_id"`
	Count   int                `bson:"count"`
}
