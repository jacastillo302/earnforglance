package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSitemapSettings = "sitemap_settings"
)

// SitemapSettings represents sitemap settings
type SitemapSettings struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	SitemapEnabled              bool               `bson:"sitemap_enabled"`
	SitemapPageSize             int                `bson:"sitemap_page_size"`
	SitemapIncludeBlogPosts     bool               `bson:"sitemap_include_blog_posts"`
	SitemapIncludeCategories    bool               `bson:"sitemap_include_categories"`
	SitemapIncludeManufacturers bool               `bson:"sitemap_include_manufacturers"`
	SitemapIncludeNews          bool               `bson:"sitemap_include_news"`
	SitemapIncludeProducts      bool               `bson:"sitemap_include_products"`
	SitemapIncludeProductTags   bool               `bson:"sitemap_include_product_tags"`
	SitemapIncludeTopics        bool               `bson:"sitemap_include_topics"`
}
