package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionLocalizationSettings = "localization_settings"
)

// LocalizationSettings represents localization settings all
type LocalizationSettings struct {
	ID                                  bson.ObjectID `bson:"_id,omitempty"`
	DefaultAdminLanguageID              bson.ObjectID `bson:"default_admin_language_id"`
	UseImagesForLanguageSelection       bool          `bson:"use_images_for_language_selection"`
	SeoFriendlyUrlsForLanguagesEnabled  bool          `bson:"seo_friendly_urls_for_languages_enabled"`
	AutomaticallyDetectLanguage         bool          `bson:"automatically_detect_language"`
	LoadAllLocaleRecordsOnStartup       bool          `bson:"load_all_locale_records_on_startup"`
	LoadAllLocalizedPropertiesOnStartup bool          `bson:"load_all_localized_properties_on_startup"`
	LoadAllUrlRecordsOnStartup          bool          `bson:"load_all_url_records_on_startup"`
	IgnoreRtlPropertyForAdminArea       bool          `bson:"ignore_rtl_property_for_admin_area"`
}

type LocalizationSettingsRepository interface {
	CreateMany(c context.Context, items []LocalizationSettings) error
	Create(c context.Context, localization_settings *LocalizationSettings) error
	Update(c context.Context, localization_settings *LocalizationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocalizationSettings, error)
	FetchByID(c context.Context, ID string) (LocalizationSettings, error)
}

type LocalizationSettingsUsecase interface {
	CreateMany(c context.Context, items []LocalizationSettings) error
	FetchByID(c context.Context, ID string) (LocalizationSettings, error)
	Create(c context.Context, localization_settings *LocalizationSettings) error
	Update(c context.Context, localization_settings *LocalizationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocalizationSettings, error)
}
