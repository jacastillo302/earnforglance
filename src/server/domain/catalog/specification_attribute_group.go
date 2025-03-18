package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSpecificationAttributeGroup = "specification_attribute_groups"
)

// SpecificationAttributeGroup represents a specification attribute group
type SpecificationAttributeGroup struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}

type SpecificationAttributeGroupRepository interface {
	Create(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Update(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Delete(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Fetch(c context.Context) ([]SpecificationAttributeGroup, error)
	FetchByID(c context.Context, specification_attribute_groupID string) (SpecificationAttributeGroup, error)
}

type SpecificationAttributeGroupUsecase interface {
	FetchByID(c context.Context, specification_attribute_groupID string) (SpecificationAttributeGroup, error)
	Create(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Update(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Delete(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Fetch(c context.Context) ([]SpecificationAttributeGroup, error)
}
