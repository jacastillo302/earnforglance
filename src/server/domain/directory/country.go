package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCountry = "countries"
)

// Country represents a country
type Country struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	AllowsBilling      bool               `bson:"allows_billing"`
	AllowsShipping     bool               `bson:"allows_shipping"`
	TwoLetterIsoCode   string             `bson:"two_letter_iso_code"`
	ThreeLetterIsoCode string             `bson:"three_letter_iso_code"`
	NumericIsoCode     int                `bson:"numeric_iso_code"`
	SubjectToVat       bool               `bson:"subject_to_vat"`
	Published          bool               `bson:"published"`
	DisplayOrder       int                `bson:"display_order"`
	LimitedToStores    bool               `bson:"limited_to_stores"`
}
