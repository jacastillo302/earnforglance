package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCustomer = "customers"
)

// Customer represents a customer
type Customer struct {
	ID                                 bson.ObjectID  `bson:"_id,omitempty"`
	CustomerGuid                       string         `bson:"customer_guid"`
	Username                           string         `bson:"username"`
	Email                              string         `bson:"email"`
	FirstName                          string         `bson:"first_name"`
	LastName                           string         `bson:"last_name"`
	Gender                             string         `bson:"gender"`
	DateOfBirth                        *time.Time     `bson:"date_of_birth"`
	Company                            string         `bson:"company"`
	StreetAddress                      string         `bson:"street_address"`
	StreetAddress2                     string         `bson:"street_address2"`
	ZipPostalCode                      string         `bson:"zip_postal_code"`
	City                               string         `bson:"city"`
	County                             string         `bson:"county"`
	CountryID                          bson.ObjectID  `bson:"country_id"`
	StateProvinceID                    bson.ObjectID  `bson:"state_province_id"`
	Phone                              string         `bson:"phone"`
	Fax                                string         `bson:"fax"`
	VatNumber                          string         `bson:"vat_number"`
	VatNumberStatusID                  string         `bson:"vat_number_status_id"`
	TimeZoneID                         string         `bson:"time_zone_id"`
	CustomPermisionRecordAttributesXML string         `bson:"custom_customer_attributes_xml"`
	CurrencyID                         *bson.ObjectID `bson:"currency_id"`
	LanguageID                         *bson.ObjectID `bson:"language_id"`
	TaxDisplayTypeID                   *int           `bson:"tax_display_type_id"`
	EmailToRevalidate                  string         `bson:"email_to_revalidate"`
	AdminComment                       string         `bson:"admin_comment"`
	IsTaxExempt                        bool           `bson:"is_tax_exempt"`
	AffiliateID                        string         `bson:"affiliate_id"`
	VendorID                           string         `bson:"vendor_id"`
	HasShoppingCartItems               bool           `bson:"has_shopping_cart_items"`
	RequireReLogin                     bool           `bson:"require_re_login"`
	FailedLoginAttempts                int            `bson:"failed_login_attempts"`
	CannotLoginUntilDateUtc            *time.Time     `bson:"cannot_login_until_date_utc"`
	Active                             bool           `bson:"active"`
	Deleted                            bool           `bson:"deleted"`
	IsSystemAccount                    bool           `bson:"is_system_account"`
	SystemName                         string         `bson:"system_name"`
	LastIpAddress                      string         `bson:"last_ip_address"`
	CreatedOnUtc                       time.Time      `bson:"created_on_utc"`
	LastLoginDateUtc                   *time.Time     `bson:"last_login_date_utc"`
	LastActivityDateUtc                time.Time      `bson:"last_activity_date_utc"`
	RegisteredInStoreID                string         `bson:"registered_in_store_id"`
	BillingAddressID                   *bson.ObjectID `bson:"billing_address_id"`
	MustChangePassword                 bool           `bson:"must_change_password"`
	ShippingAddressID                  *bson.ObjectID `bson:"shipping_address_id"`
}

type CustomerRepository interface {
	CreateMany(c context.Context, items []Customer) error
	Create(c context.Context, customer *Customer) error
	Update(c context.Context, customer *Customer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Customer, error)
	FetchByID(c context.Context, ID string) (Customer, error)
}

type CustomerUsecase interface {
	CreateMany(c context.Context, items []Customer) error
	FetchByID(c context.Context, ID string) (Customer, error)
	Create(c context.Context, customer *Customer) error
	Update(c context.Context, customer *Customer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Customer, error)
}

// NewCustomer creates a new Customer instance
func NewCustomer() *Customer {
	return &Customer{
		CustomerGuid: uuid.New().String(),
		CreatedOnUtc: time.Now(),
	}
}
