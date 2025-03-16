package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionStoreInformationSettings = "store_information_settings"
)

// StoreInformationSettings represents store information settings
type StoreInformationSettings struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	HidePoweredBy              bool               `bson:"hide_powered_by"`
	StoreClosed                bool               `bson:"store_closed"`
	LogoPictureID              int                `bson:"logo_picture_id"`
	DefaultStoreTheme          string             `bson:"default_store_theme"`
	AllowCustomerToSelectTheme bool               `bson:"allow_customer_to_select_theme"`
	DisplayEuCookieLawWarning  bool               `bson:"display_eu_cookie_law_warning"`
	FacebookLink               string             `bson:"facebook_link"`
	TwitterLink                string             `bson:"twitter_link"`
	YoutubeLink                string             `bson:"youtube_link"`
	InstagramLink              string             `bson:"instagram_link"`
}
