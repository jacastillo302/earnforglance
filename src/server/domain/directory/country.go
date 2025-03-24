package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type CountryRepository interface {
	CreateMany(c context.Context, items []Country) error
	Create(c context.Context, country *Country) error
	Update(c context.Context, country *Country) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Country, error)
	FetchByID(c context.Context, ID string) (Country, error)
}

type CountryUsecase interface {
	FetchByID(c context.Context, ID string) (Country, error)
	Create(c context.Context, country *Country) error
	Update(c context.Context, country *Country) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Country, error)
}
