package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionStoreMapping = "store_mappings"
)

// StoreMapping represents a store mapping record.
type StoreMapping struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	EntityID   bson.ObjectID `bson:"entity_id"`
	EntityName string        `bson:"entity_name"`
	StoreID    bson.ObjectID `bson:"store_id"`
}

// StoreMappingRepository defines the repository interface for StoreMapping
type StoreMappingRepository interface {
	CreateMany(c context.Context, items []StoreMapping) error
	Create(c context.Context, store_mapping *StoreMapping) error
	Update(c context.Context, store_mapping *StoreMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StoreMapping, error)
	FetchByID(c context.Context, ID string) (StoreMapping, error)
}

// StoreMappingUsecase defines the use case interface for StoreMapping
type StoreMappingUsecase interface {
	CreateMany(c context.Context, items []StoreMapping) error
	FetchByID(c context.Context, ID string) (StoreMapping, error)
	Create(c context.Context, store_mapping *StoreMapping) error
	Update(c context.Context, store_mapping *StoreMapping) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StoreMapping, error)
}
