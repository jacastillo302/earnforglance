package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionLocaleStringResource = "locale_string_resources"
)

// LocaleStringResource represents a locale string resource
type LocaleStringResource struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID    int                `bson:"language_id"`
	ResourceName  string             `bson:"resource_name"`
	ResourceValue string             `bson:"resource_value"`
}
