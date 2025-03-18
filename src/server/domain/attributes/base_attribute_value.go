package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBaseAttributeValue = "base_attribute_values"
)

// BaseAttributeValue represents the base class for attribute values
type BaseAttributevValue struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	IsPreSelected bool               `bson:"is_pre_selected"`
	DisplayOrder  int                `bson:"display_order"`
	AttributeId   primitive.ObjectID `bson:"attribute_id"`
}

type BaseAttributevValueRepository interface {
	Create(c context.Context, base_attribute *BaseAttributevValue) error
	Update(c context.Context, base_attribute *BaseAttributevValue) error
	Delete(c context.Context, base_attribute *BaseAttributevValue) error
	Fetch(c context.Context) ([]BaseAttributevValue, error)
	FetchByID(c context.Context, base_attributeID string) (BaseAttributevValue, error)
}

type BaseAttributevValueUsecase interface {
	FetchByID(c context.Context, base_attributeID string) (BaseAttributevValue, error)
	Create(c context.Context, base_attribute *BaseAttributevValue) error
	Update(c context.Context, base_attribute *BaseAttributevValue) error
	Delete(c context.Context, base_attribute *BaseAttributevValue) error
	Fetch(c context.Context) ([]BaseAttributevValue, error)
}
