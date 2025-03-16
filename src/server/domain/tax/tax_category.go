package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectioTaxCategory = "tax_categories"
)

// TaxCategory represents a tax category
type TaxCategory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}
