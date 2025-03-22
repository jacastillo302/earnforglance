package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCommonSettings = "common_settings"
)

// CommonSettings represents common settings
type CommonSettings struct {
	ID                               primitive.ObjectID `bson:"_id,omitempty"`
	SubjectFieldOnContactUsForm      bool               `bson:"subject_field_on_contact_us_form"`
	UseSystemEmailForContactUsForm   bool               `bson:"use_system_email_for_contact_us_form"`
	DisplayJavaScriptDisabledWarning bool               `bson:"display_javascript_disabled_warning"`
	Log404Errors                     bool               `bson:"log_404_errors"`
	BreadcrumbDelimiter              string             `bson:"breadcrumb_delimiter"`
	IgnoreLogWordlist                []string           `bson:"ignore_log_wordlist"`
	ClearLogOlderThanDays            int                `bson:"clear_log_older_than_days"`
	BbcodeEditorOpenLinksInNewWindow bool               `bson:"bbcode_editor_open_links_in_new_window"`
	PopupForTermsOfServiceLinks      bool               `bson:"popup_for_terms_of_service_links"`
	JqueryMigrateScriptLoggingActive bool               `bson:"jquery_migrate_script_logging_active"`
	UseResponseCompression           bool               `bson:"use_response_compression"`
	FaviconAndAppIconsHeadCode       string             `bson:"favicon_and_app_icons_head_code"`
	EnableHtmlMinification           bool               `bson:"enable_html_minification"`
	RestartTimeout                   *int               `bson:"restart_timeout,omitempty"`
	HeaderCustomHtml                 string             `bson:"header_custom_html"`
	FooterCustomHtml                 string             `bson:"footer_custom_html"`
}

type CommonSettingsRepository interface {
	Create(c context.Context, common_settings *CommonSettings) error
	Update(c context.Context, common_settings *CommonSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CommonSettings, error)
	FetchByID(c context.Context, ID string) (CommonSettings, error)
}

type CommonSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (CommonSettings, error)
	Create(c context.Context, common_settings *CommonSettings) error
	Update(c context.Context, common_settings *CommonSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CommonSettings, error)
}

// NewCommonSettings creates a new CommonSettings instance
func NewCommonSettings() *CommonSettings {
	return &CommonSettings{
		IgnoreLogWordlist: []string{},
	}
}
