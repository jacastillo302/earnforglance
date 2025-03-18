package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategoryTemplate = "category_templates"
)

// CategoryTemplate represents a category template
type CategoryTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	ViewPath     string             `bson:"view_path"`
	DisplayOrder int                `bson:"display_order"`
}

type CategoryTemplateRepository interface {
	Create(c context.Context, category_template *CategoryTemplate) error
	Update(c context.Context, category_template *CategoryTemplate) error
	Delete(c context.Context, category_template *CategoryTemplate) error
	Fetch(c context.Context) ([]CategoryTemplate, error)
	FetchByID(c context.Context, category_templateID string) (CategoryTemplate, error)
}

type CategoryTemplateUsecase interface {
	FetchByID(c context.Context, category_templateID string) (CategoryTemplate, error)
	Create(c context.Context, category_template *CategoryTemplate) error
	Update(c context.Context, category_template *CategoryTemplate) error
	Delete(c context.Context, category_template *CategoryTemplate) error
	Fetch(c context.Context) ([]CategoryTemplate, error)
}
