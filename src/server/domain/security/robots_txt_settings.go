package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionRobotsTxtSettings = "robots_txt_settings"
)

// RobotsTxtSettings represents robots.txt settings
type RobotsTxtSettings struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	DisallowPaths            []string           `bson:"disallow_paths"`
	LocalizableDisallowPaths []string           `bson:"localizable_disallow_paths"`
	DisallowLanguages        []int              `bson:"disallow_languages"`
	AdditionsRules           []string           `bson:"additions_rules"`
	AllowSitemapXml          bool               `bson:"allow_sitemap_xml"`
}

// NewRobotsTxtSettings creates a new instance of RobotsTxtSettings with default values
func NewRobotsTxtSettings() *RobotsTxtSettings {
	return &RobotsTxtSettings{
		DisallowPaths:            []string{},
		LocalizableDisallowPaths: []string{},
		DisallowLanguages:        []int{},
		AdditionsRules:           []string{},
		AllowSitemapXml:          true,
	}
}
