package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionStoreInformationSettings = "store_information_settings"
)

// StoreInformationSettings represents store information settings.
type StoreInformationSettings struct {
	ID                         bson.ObjectID `bson:"_id,omitempty"`
	HidePoweredBy              bool          `bson:"hide_powered_by"`
	StoreClosed                bool          `bson:"store_closed"`
	LogoPictureID              int           `bson:"logo_picture_id"`
	DefaultStoreTheme          string        `bson:"default_store_theme"`
	AllowCustomerToSelectTheme bool          `bson:"allow_customer_to_select_theme"`
	DisplayEuCookieLawWarning  bool          `bson:"display_eu_cookie_law_warning"`
	FacebookLink               string        `bson:"facebook_link"`
	TwitterLink                string        `bson:"twitter_link"`
	YoutubeLink                string        `bson:"youtube_link"`
	InstagramLink              string        `bson:"instagram_link"`
}

// StoreInformationSettingsRepository defines the repository interface for StoreInformationSettings
type StoreInformationSettingsRepository interface {
	CreateMany(c context.Context, items []StoreInformationSettings) error
	Create(c context.Context, store_information_settings *StoreInformationSettings) error
	Update(c context.Context, store_information_settings *StoreInformationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StoreInformationSettings, error)
	FetchByID(c context.Context, ID string) (StoreInformationSettings, error)
}

// StoreInformationSettingsUsecase defines the use case interface for StoreInformationSettings
type StoreInformationSettingsUsecase interface {
	CreateMany(c context.Context, items []StoreInformationSettings) error
	FetchByID(c context.Context, ID string) (StoreInformationSettings, error)
	Create(c context.Context, store_information_settings *StoreInformationSettings) error
	Update(c context.Context, store_information_settings *StoreInformationSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StoreInformationSettings, error)
}
