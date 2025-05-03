package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionBaseAttributeValue = "base_attribute_values"
)

// BaseAttributeValue represents the base class for attribute values
type BaseAttributeValue struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	BaseAttributeID bson.ObjectID `bson:"base_attribute_id"`
	Name            string        `bson:"name"`
	IsPreSelected   bool          `bson:"is_pre_selected"`
	DisplayOrder    int           `bson:"display_order"`
}

type BaseAttributeValueRepository interface {
	CreateMany(c context.Context, items []BaseAttributeValue) error
	Create(c context.Context, base_attribute *BaseAttributeValue) error
	Update(c context.Context, base_attribute *BaseAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BaseAttributeValue, error)
	FetchByID(c context.Context, ID string) (BaseAttributeValue, error)
}

type BaseAttributeValueUsecase interface {
	CreateMany(c context.Context, items []BaseAttributeValue) error
	FetchByID(c context.Context, ID string) (BaseAttributeValue, error)
	Create(c context.Context, base_attribute *BaseAttributeValue) error
	Update(c context.Context, base_attribute *BaseAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BaseAttributeValue, error)
}
