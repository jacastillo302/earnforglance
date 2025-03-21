package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStoreMapping = "store_mappings"
)

// StoreMapping represents a store mapping record.
type StoreMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EntityID   primitive.ObjectID `bson:"entity_id"`
	EntityName string             `bson:"entity_name"`
	StoreID    primitive.ObjectID `bson:"store_id"`
}

// StoreMappingRepository defines the repository interface for StoreMapping
type StoreMappingRepository interface {
	Create(c context.Context, store_mapping *StoreMapping) error
	Update(c context.Context, store_mapping *StoreMapping) error
	Delete(c context.Context, store_mapping *StoreMapping) error
	Fetch(c context.Context) ([]StoreMapping, error)
	FetchByID(c context.Context, store_mappingID string) (StoreMapping, error)
}

// StoreMappingUsecase defines the use case interface for StoreMapping
type StoreMappingUsecase interface {
	FetchByID(c context.Context, store_mappingID string) (StoreMapping, error)
	Create(c context.Context, store_mapping *StoreMapping) error
	Update(c context.Context, store_mapping *StoreMapping) error
	Delete(c context.Context, store_mapping *StoreMapping) error
	Fetch(c context.Context) ([]StoreMapping, error)
}
