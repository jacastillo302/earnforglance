package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTaxCategory = "tax_categories"
)

// TaxCategory represents a tax category.
type TaxCategory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}

type TaxCategoryRepository interface {
	Create(c context.Context, tax_category *TaxCategory) error
	Update(c context.Context, tax_category *TaxCategory) error
	Delete(c context.Context, tax_category *TaxCategory) error
	Fetch(c context.Context) ([]TaxCategory, error)
	FetchByID(c context.Context, tax_categoryID string) (TaxCategory, error)
}

type TaxCategoryUsecase interface {
	FetchByID(c context.Context, tax_categoryID string) (TaxCategory, error)
	Create(c context.Context, tax_category *TaxCategory) error
	Update(c context.Context, tax_category *TaxCategory) error
	Delete(c context.Context, tax_category *TaxCategory) error
	Fetch(c context.Context) ([]TaxCategory, error)
}
