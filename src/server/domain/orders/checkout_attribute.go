package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCheckoutAttributem = "checkout_attributes"
)

// CheckoutAttribute represents a checkout attribute
type CheckoutAttribute struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty"`
	TextPrompt                      string             `bson:"text_prompt"`
	ShippableProductRequired        bool               `bson:"shippable_product_required"`
	IsTaxExempt                     bool               `bson:"is_tax_exempt"`
	TaxCategoryID                   int                `bson:"tax_category_id"`
	LimitedToStores                 bool               `bson:"limited_to_stores"`
	ValidationMinLength             *int               `bson:"validation_min_length,omitempty"`
	ValidationMaxLength             *int               `bson:"validation_max_length,omitempty"`
	ValidationFileAllowedExtensions string             `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int               `bson:"validation_file_maximum_size,omitempty"`
	DefaultValue                    string             `bson:"default_value"`
	ConditionAttributeXml           string             `bson:"condition_attribute_xml"`
}
