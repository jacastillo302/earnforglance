package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionAddress = "addresses"
)

// Address represents an address
type Address struct {
	ID               bson.ObjectID  `bson:"_id,omitempty"`
	FirstName        string         `bson:"first_name"`
	LastName         string         `bson:"last_name"`
	Email            string         `bson:"email"`
	Company          string         `bson:"company"`
	CountryID        *bson.ObjectID `bson:"country_id"`
	StateProvinceID  *bson.ObjectID `bson:"state_province_id"`
	County           string         `bson:"county"`
	City             string         `bson:"city"`
	Address1         string         `bson:"address1"`
	Address2         string         `bson:"address2"`
	ZipPostalCode    string         `bson:"zip_postal_code"`
	PhoneNumber      string         `bson:"phone_number"`
	FaxNumber        string         `bson:"fax_number"`
	CustomAttributes string         `bson:"custom_attributes"`
	CreatedOnUtc     time.Time      `bson:"created_on_utc"`
}

type AddressRepository interface {
	CreateMany(c context.Context, items []Address) error
	Create(c context.Context, product_tag *Address) error
	Update(c context.Context, product_tag *Address) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Address, error)
	FetchByID(c context.Context, ID string) (Address, error)
}

type AddressUsecase interface {
	CreateMany(c context.Context, items []Address) error
	FetchByID(c context.Context, ID string) (Address, error)
	Create(c context.Context, product_tag *Address) error
	Update(c context.Context, product_tag *Address) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Address, error)
}
