package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSpecificationAttribute = "specification_attributes"
)

// SpecificationAttribute represents a specification attribute
type SpecificationAttribute struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`
	Name                          string             `bson:"name"`
	DisplayOrder                  int                `bson:"display_order"`
	SpecificationAttributeGroupID *int               `bson:"specification_attribute_group_id"`
}

type SpecificationAttributeRepository interface {
	CreateMany(c context.Context, items []SpecificationAttribute) error
	Create(c context.Context, specification_attribute *SpecificationAttribute) error
	Update(c context.Context, specification_attribute *SpecificationAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttribute, error)
	FetchByID(c context.Context, ID string) (SpecificationAttribute, error)
}

type SpecificationAttributeUsecase interface {
	CreateMany(c context.Context, items []SpecificationAttribute) error
	FetchByID(c context.Context, ID string) (SpecificationAttribute, error)
	Create(c context.Context, specification_attribute *SpecificationAttribute) error
	Update(c context.Context, specification_attribute *SpecificationAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttribute, error)
}
