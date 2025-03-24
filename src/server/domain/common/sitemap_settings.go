package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSitemapSettings = "sitemap_settings"
)

// SitemapSettings represents sitemap settings
type SitemapSettings struct {
	ID                             primitive.ObjectID `bson:"_id,omitempty"`
	SitemapEnabled                 bool               `bson:"sitemap_enabled"`
	SitemapPageSize                int                `bson:"sitemap_page_size"`
	SitemapIncludeBlogPosts        bool               `bson:"sitemap_include_blog_posts"`
	SitemapIncludeCategories       bool               `bson:"sitemap_include_categories"`
	SitemapIncludeManufacturers    bool               `bson:"sitemap_include_manufacturers"`
	SitemapIncludeNews             bool               `bson:"sitemap_include_news"`
	SitemapIncludeProducts         bool               `bson:"sitemap_include_products"`
	SitemapIncludeSitemapSettingss bool               `bson:"sitemap_include_sitemap_settingss"`
	SitemapIncludeTopics           bool               `bson:"sitemap_include_topics"`
}

type SitemapSettingsRepository interface {
	CreateMany(c context.Context, items []SitemapSettings) error
	Create(c context.Context, sitemap_settings *SitemapSettings) error
	Update(c context.Context, sitemap_settings *SitemapSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SitemapSettings, error)
	FetchByID(c context.Context, ID string) (SitemapSettings, error)
}

type SitemapSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (SitemapSettings, error)
	Create(c context.Context, sitemap_settings *SitemapSettings) error
	Update(c context.Context, sitemap_settings *SitemapSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]SitemapSettings, error)
}
