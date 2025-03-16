package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionLocalizationSettings = "localization_settings"
)

// LocalizationSettings represents localization settings
type LocalizationSettings struct {
	ID                                  primitive.ObjectID `bson:"_id,omitempty"`
	DefaultAdminLanguageID              int                `bson:"default_admin_language_id"`
	UseImagesForLanguageSelection       bool               `bson:"use_images_for_language_selection"`
	SeoFriendlyUrlsForLanguagesEnabled  bool               `bson:"seo_friendly_urls_for_languages_enabled"`
	AutomaticallyDetectLanguage         bool               `bson:"automatically_detect_language"`
	LoadAllLocaleRecordsOnStartup       bool               `bson:"load_all_locale_records_on_startup"`
	LoadAllLocalizedPropertiesOnStartup bool               `bson:"load_all_localized_properties_on_startup"`
	LoadAllUrlRecordsOnStartup          bool               `bson:"load_all_url_records_on_startup"`
	IgnoreRtlPropertyForAdminArea       bool               `bson:"ignore_rtl_property_for_admin_area"`
}
