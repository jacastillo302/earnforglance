package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionRobotsTxtSettings = "robots_txt_settings"
)

// RobotsTxtSettings represents robots.txt settings
type RobotsTxtSettings struct {
	ID                       bson.ObjectID `bson:"_id,omitempty"`
	DisallowPaths            []string      `bson:"disallow_paths"`
	LocalizableDisallowPaths []string      `bson:"localizable_disallow_paths"`
	DisallowLanguages        []int         `bson:"disallow_languages"`
	AdditionsRules           []string      `bson:"additions_rules"`
	AllowSitemapXml          bool          `bson:"allow_sitemap_xml"`
}

// NewRobotsTxtSettings creates a new instance of RobotsTxtSettings with default values.
func NewRobotsTxtSettings() *RobotsTxtSettings {
	return &RobotsTxtSettings{
		DisallowPaths:            []string{},
		LocalizableDisallowPaths: []string{},
		DisallowLanguages:        []int{},
		AdditionsRules:           []string{},
		AllowSitemapXml:          true,
	}
}

// RobotsTxtSettingsRepository defines the repository interface for RobotsTxtSettings
type RobotsTxtSettingsRepository interface {
	CreateMany(c context.Context, items []RobotsTxtSettings) error
	Create(c context.Context, robotsTxtSettings *RobotsTxtSettings) error
	Update(c context.Context, robotsTxtSettings *RobotsTxtSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RobotsTxtSettings, error)
	FetchByID(c context.Context, ID string) (RobotsTxtSettings, error)
}

// RobotsTxtSettingsUsecase defines the use case interface for RobotsTxtSettings
type RobotsTxtSettingsUsecase interface {
	CreateMany(c context.Context, items []RobotsTxtSettings) error
	FetchByID(c context.Context, ID string) (RobotsTxtSettings, error)
	Create(c context.Context, robotsTxtSettings *RobotsTxtSettings) error
	Update(c context.Context, robotsTxtSettings *RobotsTxtSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RobotsTxtSettings, error)
}
