package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAddressSettings = "address_settings"
)

// AddressSettings represents address settings
type AddressSettings struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`
	CompanyEnabled            bool                `bson:"company_enabled"`
	CompanyRequired           bool                `bson:"company_required"`
	StreetAddressEnabled      bool                `bson:"street_address_enabled"`
	StreetAddressRequired     bool                `bson:"street_address_required"`
	StreetAddress2Enabled     bool                `bson:"street_address2_enabled"`
	StreetAddress2Required    bool                `bson:"street_address2_required"`
	ZipPostalCodeEnabled      bool                `bson:"zip_postal_code_enabled"`
	ZipPostalCodeRequired     bool                `bson:"zip_postal_code_required"`
	CityEnabled               bool                `bson:"city_enabled"`
	CityRequired              bool                `bson:"city_required"`
	CountyEnabled             bool                `bson:"county_enabled"`
	CountyRequired            bool                `bson:"county_required"`
	CountryEnabled            bool                `bson:"country_enabled"`
	DefaultCountryID          *primitive.ObjectID `bson:"default_country_id,omitempty"`
	StateProvinceEnabled      bool                `bson:"state_province_enabled"`
	PhoneEnabled              bool                `bson:"phone_enabled"`
	PhoneRequired             bool                `bson:"phone_required"`
	FaxEnabled                bool                `bson:"fax_enabled"`
	FaxRequired               bool                `bson:"fax_required"`
	PreselectCountryIfOnlyOne bool                `bson:"preselect_country_if_only_one"`
}

type AddressSettingsRepository interface {
	CreateMany(c context.Context, items []AddressSettings) error
	Create(c context.Context, address_settings *AddressSettings) error
	Update(c context.Context, address_settings *AddressSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AddressSettings, error)
	FetchByID(c context.Context, ID string) (AddressSettings, error)
}

type AddressSettingsUsecase interface {
	CreateMany(c context.Context, items []AddressSettings) error
	FetchByID(c context.Context, ID string) (AddressSettings, error)
	Create(c context.Context, address_settings *AddressSettings) error
	Update(c context.Context, address_settings *AddressSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AddressSettings, error)
}
