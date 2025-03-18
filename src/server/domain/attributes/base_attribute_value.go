package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBaseAttributeValue = "base_attribute_values"
)

// BaseAttributeValue represents the base class for attribute values
type BaseAttributeValue struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	IsPreSelected bool               `bson:"is_pre_selected"`
	DisplayOrder  int                `bson:"display_order"`
	AttributeId   primitive.ObjectID `bson:"attribute_id"`
}

type BaseAttributeValueRepository interface {
	Create(c context.Context, base_attribute *BaseAttributeValue) error
	Update(c context.Context, base_attribute *BaseAttributeValue) error
	Delete(c context.Context, base_attribute *BaseAttributeValue) error
	Fetch(c context.Context) ([]BaseAttributeValue, error)
	FetchByID(c context.Context, base_attributeID string) (BaseAttributeValue, error)
}

type BaseAttributeValueUsecase interface {
	FetchByID(c context.Context, base_attributeID string) (BaseAttributeValue, error)
	Create(c context.Context, base_attribute *BaseAttributeValue) error
	Update(c context.Context, base_attribute *BaseAttributeValue) error
	Delete(c context.Context, base_attribute *BaseAttributeValue) error
	Fetch(c context.Context) ([]BaseAttributeValue, error)
}
