package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSeoSettings = "seo_settings"
)

// SeoSettings represents SEO settings.
type SeoSettings struct {
	ID                                bson.ObjectID `bson:"_id,omitempty"`
	PageTitleSeparator                string        `bson:"page_title_separator"`
	PageTitleSeoAdjustmentID          int           `bson:"page_title_seo_adjustment"`
	GenerateProductMetaDescription    bool          `bson:"generate_product_meta_description"`
	ConvertNonWesternChars            bool          `bson:"convert_non_western_chars"`
	AllowUnicodeCharsInUrls           bool          `bson:"allow_unicode_chars_in_urls"`
	CanonicalUrlsEnabled              bool          `bson:"canonical_urls_enabled"`
	QueryStringInCanonicalUrlsEnabled bool          `bson:"query_string_in_canonical_urls_enabled"`
	WwwRequirementID                  int           `bson:"www_requirement"`
	TwitterMetaTags                   bool          `bson:"twitter_meta_tags"`
	OpenGraphMetaTags                 bool          `bson:"open_graph_meta_tags"`
	ReservedUrlRecordSlugs            []string      `bson:"reserved_url_record_slugs"`
	CustomHeadTags                    string        `bson:"custom_head_tags"`
	MicrodataEnabled                  bool          `bson:"microdata_enabled"`
}

// NewSeoSettings creates a new instance of SeoSettings with default values
func NewSeoSettings() *SeoSettings {
	return &SeoSettings{
		ReservedUrlRecordSlugs: []string{},
	}
}

type SeoSettingsRepository interface {
	CreateMany(c context.Context, items []SeoSettings) error
	Create(c context.Context, seo_settings *SeoSettings) error
	Update(c context.Context, seo_settings *SeoSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SeoSettings, error)
	FetchByID(c context.Context, ID string) (SeoSettings, error)
}

type SeoSettingsUsecase interface {
	CreateMany(c context.Context, items []SeoSettings) error
	FetchByID(c context.Context, ID string) (SeoSettings, error)
	Create(c context.Context, seo_settings *SeoSettings) error
	Update(c context.Context, seo_settings *SeoSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SeoSettings, error)
}
