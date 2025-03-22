package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSearchTerm = "search_terms"
)

// SearchTerm represents a search term record (for statistics)
type SearchTerm struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Keyword string             `bson:"keyword"`
	StoreID primitive.ObjectID `bson:"store_id"`
	Count   int                `bson:"count"`
}

type SearchTermRepository interface {
	Create(c context.Context, search_term *SearchTerm) error
	Update(c context.Context, search_term *SearchTerm) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SearchTerm, error)
	FetchByID(c context.Context, ID string) (SearchTerm, error)
}

type SearchTermUsecase interface {
	FetchByID(c context.Context, ID string) (SearchTerm, error)
	Create(c context.Context, search_term *SearchTerm) error
	Update(c context.Context, search_term *SearchTerm) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SearchTerm, error)
}
