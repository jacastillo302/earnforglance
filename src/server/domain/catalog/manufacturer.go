package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionManufacturer = "manufacturers"
)

// Manufacturer represents a manufacturer
type Manufacturer struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	Name                           string             `bson:"name"`
	Description                    string             `bson:"description"`
	ManufacturerID                 primitive.ObjectID `bson:"manufacturer_id"`
	MetaKeywords                   string             `bson:"meta_keywords"`
	MetaDescription                string             `bson:"meta_description"`
	MetaTitle                      string             `bson:"meta_title"`
	PictureID                      primitive.ObjectID `bson:"picture_id"`
	PageSize                       int                `bson:"page_size"`
	AllowCustomersToSelectPageSize bool               `bson:"allow_customers_to_select_page_size"`
	PageSizeOptions                string             `bson:"page_size_options"`
	SubjectToAcl                   bool               `bson:"subject_to_acl"`
	LimitedToStores                bool               `bson:"limited_to_stores"`
	Published                      bool               `bson:"published"`
	Deleted                        bool               `bson:"deleted"`
	DisplayOrder                   int                `bson:"display_order"`
	CreatedOnUtc                   time.Time          `bson:"created_on_utc"`
	UpdatedOnUtc                   time.Time          `bson:"updated_on_utc"`
	PriceRangeFiltering            bool               `bson:"price_range_filtering"`
	PriceFrom                      float64            `bson:"price_from"`
	PriceTo                        float64            `bson:"price_to"`
	ManuallyPriceRange             bool               `bson:"manually_price_range"`
}

type ManufacturerRepository interface {
	CreateMany(c context.Context, items []Manufacturer) error
	Create(c context.Context, manufacturer *Manufacturer) error
	Update(c context.Context, manufacturer *Manufacturer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Manufacturer, error)
	FetchByID(c context.Context, ID string) (Manufacturer, error)
}

type ManufacturerUsecase interface {
	CreateMany(c context.Context, items []Manufacturer) error
	FetchByID(c context.Context, ID string) (Manufacturer, error)
	Create(c context.Context, manufacturer *Manufacturer) error
	Update(c context.Context, manufacturer *Manufacturer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Manufacturer, error)
}
