package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionLanguage = "languages"
)

// Language represents a language
type Language struct {
	ID                bson.ObjectID `bson:"_id,omitempty"`
	Name              string        `bson:"name"`
	LanguageCulture   string        `bson:"language_culture"`
	UniqueSeoCode     string        `bson:"unique_seo_code"`
	FlagImageFileName string        `bson:"flag_image_file_name"`
	Rtl               bool          `bson:"rtl"`
	LimitedToStores   bool          `bson:"limited_to_stores"`
	DefaultCurrencyID bson.ObjectID `bson:"default_currency_id"`
	Published         bool          `bson:"published"`
	DisplayOrder      int           `bson:"display_order"`
}

type LanguageRepository interface {
	CreateMany(c context.Context, items []Language) error
	Create(c context.Context, language *Language) error
	Update(c context.Context, language *Language) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Language, error)
	FetchByID(c context.Context, ID string) (Language, error)
}

type LanguageUsecase interface {
	CreateMany(c context.Context, items []Language) error
	FetchByID(c context.Context, ID string) (Language, error)
	Create(c context.Context, language *Language) error
	Update(c context.Context, language *Language) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Language, error)
}
