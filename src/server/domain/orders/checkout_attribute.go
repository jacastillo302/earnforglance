package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCheckoutAttribute = "checkout_attributes"
)

// CheckoutAttribute represents a checkout attribute
type CheckoutAttribute struct {
	ID                              bson.ObjectID `bson:"_id,omitempty"`
	Name                            string        `bson:"name"`
	IsRequired                      bool          `bson:"is_required"`
	DisplayOrder                    int           `bson:"display_order"`
	TextPrompt                      string        `bson:"text_prompt"`
	AttributeControlTypeID          int           `bson:"attribute_control_type_id"`
	ShippableProductRequired        bool          `bson:"shippable_product_required"`
	IsTaxExempt                     bool          `bson:"is_tax_exempt"`
	TaxCategoryID                   bson.ObjectID `bson:"tax_category_id"`
	LimitedToStores                 bool          `bson:"limited_to_stores"`
	ValidationMinLength             *int          `bson:"validation_min_length"`
	ValidationMaxLength             *int          `bson:"validation_max_length"`
	ValidationFileAllowedExtensions string        `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int          `bson:"validation_file_maximum_size"`
	DefaultValue                    string        `bson:"default_value"`
	ConditionAttributeXml           string        `bson:"condition_attribute_xml"`
}

// CheckoutAttributeRepository interface
type CheckoutAttributeRepository interface {
	CreateMany(c context.Context, items []CheckoutAttribute) error
	Create(c context.Context, checkout_attribute *CheckoutAttribute) error
	Update(c context.Context, checkout_attribute *CheckoutAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CheckoutAttribute, error)
	FetchByID(c context.Context, ID string) (CheckoutAttribute, error)
}

// CheckoutAttributeUsecase interface
type CheckoutAttributeUsecase interface {
	CreateMany(c context.Context, items []CheckoutAttribute) error
	FetchByID(c context.Context, ID string) (CheckoutAttribute, error)
	Create(c context.Context, checkout_attribute *CheckoutAttribute) error
	Update(c context.Context, checkout_attribute *CheckoutAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CheckoutAttribute, error)
}
