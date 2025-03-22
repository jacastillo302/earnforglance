package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSitemapXmlSettings = "sitemap_xml_settings"
)

// SitemapXmlSettings represents sitemap.xml settings
type SitemapXmlSettings struct {
	ID                                   primitive.ObjectID `bson:"_id,omitempty"`
	SitemapXmlEnabled                    bool               `bson:"sitemap_xml_enabled"`
	SitemapXmlIncludeBlogPosts           bool               `bson:"sitemap_xml_include_blog_posts"`
	SitemapXmlIncludeCategories          bool               `bson:"sitemap_xml_include_categories"`
	SitemapXmlIncludeCustomUrls          bool               `bson:"sitemap_xml_include_custom_urls"`
	SitemapXmlIncludeManufacturers       bool               `bson:"sitemap_xml_include_manufacturers"`
	SitemapXmlIncludeNews                bool               `bson:"sitemap_xml_include_news"`
	SitemapXmlIncludeProducts            bool               `bson:"sitemap_xml_include_products"`
	SitemapXmlIncludeSitemapXmlSettingss bool               `bson:"sitemap_xml_include_sitemap_xml_settingss"`
	SitemapXmlIncludeTopics              bool               `bson:"sitemap_xml_include_topics"`
	SitemapCustomUrls                    []string           `bson:"sitemap_custom_urls"`
	RebuildSitemapXmlAfterHours          int                `bson:"rebuild_sitemap_xml_after_hours"`
	SitemapBuildOperationDelay           int                `bson:"sitemap_build_operation_delay"`
}

type SitemapXmlSettingsRepository interface {
	Create(c context.Context, sitemap_xml_settings *SitemapXmlSettings) error
	Update(c context.Context, sitemap_xml_settings *SitemapXmlSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SitemapXmlSettings, error)
	FetchByID(c context.Context, ID string) (SitemapXmlSettings, error)
}

type SitemapXmlSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (SitemapXmlSettings, error)
	Create(c context.Context, sitemap_xml_settings *SitemapXmlSettings) error
	Update(c context.Context, sitemap_xml_settings *SitemapXmlSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SitemapXmlSettings, error)
}

// NewSitemapXmlSettings creates a new SitemapXmlSettings instance
func NewSitemapXmlSettings() *SitemapXmlSettings {
	return &SitemapXmlSettings{
		SitemapCustomUrls: []string{},
	}
}
