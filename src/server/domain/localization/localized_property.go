package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionLocalizedProperty = "localized_properties"
)

// LocalizedProperty represents a localized property
type LocalizedProperty struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	EntityID       primitive.ObjectID `bson:"entity_id"`
	LanguageID     primitive.ObjectID `bson:"language_id"`
	LocaleKeyGroup string             `bson:"locale_key_group"`
	LocaleKey      string             `bson:"locale_key"`
	LocaleValue    string             `bson:"locale_value"`
}
