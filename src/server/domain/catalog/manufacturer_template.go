package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionManufacturerTemplate = "manufacturer_templates"
)

// ManufacturerTemplate represents a manufacturer template
type ManufacturerTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}

type ManufacturerTemplateRepository interface {
	Create(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Update(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Delete(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Fetch(c context.Context) ([]ManufacturerTemplate, error)
	FetchByID(c context.Context, manufacturer_templateID string) (ManufacturerTemplate, error)
}

type ManufacturerTemplateUsecase interface {
	FetchByID(c context.Context, manufacturer_templateID string) (ManufacturerTemplate, error)
	Create(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Update(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Delete(c context.Context, manufacturer_template *ManufacturerTemplate) error
	Fetch(c context.Context) ([]ManufacturerTemplate, error)
}
