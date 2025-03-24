package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAdminAreaSettings = "admin_area_settings"
)

// AdminAreaSettings represents admin area settings
type AdminAreaSettings struct {
	ID                              primitive.ObjectID `bson:"_id,omitempty"`
	DefaultGridPageSize             int                `bson:"default_grid_page_size"`
	ProductsBulkEditGridPageSize    int                `bson:"products_bulk_edit_grid_page_size"`
	PopupGridPageSize               int                `bson:"popup_grid_page_size"`
	GridPageSizes                   string             `bson:"grid_page_sizes"`
	RichEditorAdditionalSettings    string             `bson:"rich_editor_additional_settings"`
	RichEditorAllowJavaScript       bool               `bson:"rich_editor_allow_javascript"`
	RichEditorAllowStyleTag         bool               `bson:"rich_editor_allow_style_tag"`
	UseRichEditorForCustomerEmails  bool               `bson:"use_rich_editor_for_customer_emails"`
	UseRichEditorInMessageTemplates bool               `bson:"use_rich_editor_in_message_templates"`
	HideAdvertisementsOnAdminArea   bool               `bson:"hide_advertisements_on_admin_area"`
	CheckLicense                    bool               `bson:"check_license"`
	LastNewsTitleAdminArea          string             `bson:"last_news_title_admin_area"`
	UseIsoDateFormatInJsonResult    bool               `bson:"use_iso_date_format_in_json_result"`
	ShowDocumentationReferenceLinks bool               `bson:"show_documentation_reference_links"`
	UseStickyHeaderLayout           bool               `bson:"use_sticky_header_layout"`
	MinimumDropdownItemsForSearch   int                `bson:"minimum_dropdown_items_for_search"`
}

type AdminAreaSettingsRepository interface {
	CreateMany(c context.Context, items []AdminAreaSettings) error
	Create(c context.Context, admin_area_settings *AdminAreaSettings) error
	Update(c context.Context, admin_area_settings *AdminAreaSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AdminAreaSettings, error)
	FetchByID(c context.Context, ID string) (AdminAreaSettings, error)
}

type AdminAreaSettingsUsecase interface {
	CreateMany(c context.Context, items []AdminAreaSettings) error
	FetchByID(c context.Context, ID string) (AdminAreaSettings, error)
	Create(c context.Context, admin_area_settings *AdminAreaSettings) error
	Update(c context.Context, admin_area_settings *AdminAreaSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]AdminAreaSettings, error)
}
