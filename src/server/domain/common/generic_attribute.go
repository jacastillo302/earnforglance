package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGenericAttribute = "generic_attributes"
)

// GenericAttribute represents a generic attribute
type GenericAttribute struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	EntityID                primitive.ObjectID `bson:"entity_id"`
	KeyGroup                string             `bson:"key_group"`
	Key                     string             `bson:"key"`
	Value                   string             `bson:"value"`
	StoreID                 primitive.ObjectID `bson:"store_id"`
	CreatedOrUpdatedDateUTC *time.Time         `bson:"created_or_updated_date_utc,omitempty"`
}

type GenericAttributeRepository interface {
	CreateMany(c context.Context, items []GenericAttribute) error
	Create(c context.Context, generic_attribute *GenericAttribute) error
	Update(c context.Context, generic_attribute *GenericAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GenericAttribute, error)
	FetchByID(c context.Context, ID string) (GenericAttribute, error)
}

type GenericAttributeUsecase interface {
	CreateMany(c context.Context, items []GenericAttribute) error
	FetchByID(c context.Context, ID string) (GenericAttribute, error)
	Create(c context.Context, generic_attribute *GenericAttribute) error
	Update(c context.Context, generic_attribute *GenericAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GenericAttribute, error)
}
