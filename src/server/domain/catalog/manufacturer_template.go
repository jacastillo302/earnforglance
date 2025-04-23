package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionManufacturerTemplate = "manufacturer_templates"
)

// ManufacturerTemplate represents a manufacturer template
type ManufacturerTemplate struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	ViewPath     string        `bson:"view_path"`
	DisplayOrder int           `bson:"display_order"`
}

type ManufacturerTemplateRepository interface {
	CreateMany(c context.Context, items []ManufacturerTemplate) error
	Create(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Update(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ManufacturerTemplate, error)
	FetchByID(c context.Context, ID string) (ManufacturerTemplate, error)
}

type ManufacturerTemplateUsecase interface {
	CreateMany(c context.Context, items []ManufacturerTemplate) error
	FetchByID(c context.Context, ID string) (ManufacturerTemplate, error)
	Create(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Update(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ManufacturerTemplate, error)
}
