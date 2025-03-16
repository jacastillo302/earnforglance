package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSeoSettings = "seo_settings"
)

// SeoSettings represents SEO settings
type SeoSettings struct {
	ID                                primitive.ObjectID     `bson:"_id,omitempty"`
	PageTitleSeparator                string                 `bson:"page_title_separator"`
	PageTitleSeoAdjustment            PageTitleSeoAdjustment `bson:"page_title_seo_adjustment"`
	GenerateProductMetaDescription    bool                   `bson:"generate_product_meta_description"`
	ConvertNonWesternChars            bool                   `bson:"convert_non_western_chars"`
	AllowUnicodeCharsInUrls           bool                   `bson:"allow_unicode_chars_in_urls"`
	CanonicalUrlsEnabled              bool                   `bson:"canonical_urls_enabled"`
	QueryStringInCanonicalUrlsEnabled bool                   `bson:"query_string_in_canonical_urls_enabled"`
	WwwRequirement                    WwwRequirement         `bson:"www_requirement"`
	TwitterMetaTags                   bool                   `bson:"twitter_meta_tags"`
	OpenGraphMetaTags                 bool                   `bson:"open_graph_meta_tags"`
	ReservedUrlRecordSlugs            []string               `bson:"reserved_url_record_slugs"`
	CustomHeadTags                    string                 `bson:"custom_head_tags"`
	MicrodataEnabled                  bool                   `bson:"microdata_enabled"`
}

// NewSeoSettings creates a new instance of SeoSettings with default values
func NewSeoSettings() *SeoSettings {
	return &SeoSettings{
		ReservedUrlRecordSlugs: []string{},
	}
}
