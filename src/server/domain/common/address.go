package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAddress = "addresses"
)

// Address represents an address
type Address struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	FirstName        string             `bson:"first_name"`
	LastName         string             `bson:"last_name"`
	Email            string             `bson:"email"`
	Company          string             `bson:"company"`
	CountryID        *int               `bson:"country_id,omitempty"`
	StateProvinceID  *int               `bson:"state_province_id,omitempty"`
	County           string             `bson:"county"`
	City             string             `bson:"city"`
	Address1         string             `bson:"address1"`
	Address2         string             `bson:"address2"`
	ZipPostalCode    string             `bson:"zip_postal_code"`
	PhoneNumber      string             `bson:"phone_number"`
	FaxNumber        string             `bson:"fax_number"`
	CustomAttributes string             `bson:"custom_attributes"`
	CreatedOnUtc     time.Time          `bson:"created_on_utc"`
}
