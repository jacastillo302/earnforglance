package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectioStoreMapping = "store_mappings"
)

// StoreMapping represents a store mapping record
type StoreMapping struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EntityID   int                `bson:"entity_id"`
	EntityName string             `bson:"entity_name"`
	StoreID    int                `bson:"store_id"`
}
