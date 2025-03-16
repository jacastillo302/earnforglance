package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionLanguage = "languages"
)

// Language represents a language
type Language struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	LanguageCulture   string             `bson:"language_culture"`
	UniqueSeoCode     string             `bson:"unique_seo_code"`
	FlagImageFileName string             `bson:"flag_image_file_name"`
	Rtl               bool               `bson:"rtl"`
	LimitedToStores   bool               `bson:"limited_to_stores"`
	DefaultCurrencyID int                `bson:"default_currency_id"`
	Published         bool               `bson:"published"`
	DisplayOrder      int                `bson:"display_order"`
}
