package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSpecificationAttributeOption = "specification_attribute_options"
)

// SpecificationAttributeOption represents a specification attribute option
type SpecificationAttributeOption struct {
	ID                       bson.ObjectID `bson:"_id,omitempty"`
	SpecificationAttributeID bson.ObjectID `bson:"specification_attribute_id"`
	Name                     string        `bson:"name"`
	ColorSquaresRgb          string        `bson:"color_squares_rgb"`
	DisplayOrder             int           `bson:"display_order"`
}

type SpecificationAttributeOptionRepository interface {
	CreateMany(c context.Context, items []SpecificationAttributeOption) error
	Create(c context.Context, specification_attribute_option *SpecificationAttributeOption) error
	Update(c context.Context, specification_attribute_option *SpecificationAttributeOption) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttributeOption, error)
	FetchByID(c context.Context, ID string) (SpecificationAttributeOption, error)
}

type SpecificationAttributeOptionUsecase interface {
	CreateMany(c context.Context, items []SpecificationAttributeOption) error
	FetchByID(c context.Context, ID string) (SpecificationAttributeOption, error)
	Create(c context.Context, specification_attribute_option *SpecificationAttributeOption) error
	Update(c context.Context, specification_attribute_option *SpecificationAttributeOption) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SpecificationAttributeOption, error)
}
