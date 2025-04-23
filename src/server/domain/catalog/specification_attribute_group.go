package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSpecificationAttributeGroup = "specification_attribute_groups"
)

// SpecificationAttributeGroup represents a specification attribute group
type SpecificationAttributeGroup struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	DisplayOrder int           `bson:"display_order"`
}

type SpecificationAttributeGroupRepository interface {
	CreateMany(c context.Context, items []SpecificationAttributeGroup) error
	Create(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Update(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttributeGroup, error)
	FetchByID(c context.Context, ID string) (SpecificationAttributeGroup, error)
}

type SpecificationAttributeGroupUsecase interface {
	CreateMany(c context.Context, items []SpecificationAttributeGroup) error
	FetchByID(c context.Context, ID string) (SpecificationAttributeGroup, error)
	Create(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Update(c context.Context, specification_attribute_group *SpecificationAttributeGroup) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttributeGroup, error)
}
