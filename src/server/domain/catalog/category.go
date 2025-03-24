package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategory = "categories"
)

// Category represents a category
type Category struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	Name                           string             `bson:"name"`
	Description                    string             `bson:"description"`
	CategoryTemplateID             primitive.ObjectID `bson:"category_template_id"`
	MetaKeywords                   string             `bson:"meta_keywords"`
	MetaDescription                string             `bson:"meta_description"`
	MetaTitle                      string             `bson:"meta_title"`
	ParentCategoryID               primitive.ObjectID `bson:"parent_category_id"`
	PictureID                      primitive.ObjectID `bson:"picture_id"`
	PageSize                       int                `bson:"page_size"`
	AllowCustomersToSelectPageSize bool               `bson:"allow_customers_to_select_page_size"`
	PageSizeOptions                string             `bson:"page_size_options"`
	ShowOnHomepage                 bool               `bson:"show_on_homepage"`
	IncludeInTopMenu               bool               `bson:"include_in_top_menu"`
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
	RestrictFromVendors            bool               `bson:"restrict_from_vendors"`
}

type CategoryRepository interface {
	CreateMany(c context.Context, items []Category) error
	Create(c context.Context, category *Category) error
	Update(c context.Context, category *Category) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Category, error)
	FetchByID(c context.Context, ID string) (Category, error)
}

type CategoryUsecase interface {
	FetchByID(c context.Context, ID string) (Category, error)
	Create(c context.Context, category *Category) error
	Update(c context.Context, category *Category) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Category, error)
}
