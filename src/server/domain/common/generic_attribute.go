package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGenericAttribute = "generic_attributes"
)

// GenericAttribute represents a generic attribute
type GenericAttribute struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	EntityID                int                `bson:"entity_id"`
	KeyGroup                string             `bson:"key_group"`
	Key                     string             `bson:"key"`
	Value                   string             `bson:"value"`
	StoreID                 int                `bson:"store_id"`
	CreatedOrUpdatedDateUTC *time.Time         `bson:"created_or_updated_date_utc,omitempty"`
}
