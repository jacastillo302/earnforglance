package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionLocalizedProperty = "localized_properties"
)

// LocalizedProperty represents a localized property
type LocalizedProperty struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	PermissionRecordID bson.ObjectID `bson:"entity_id"`
	LanguageID         bson.ObjectID `bson:"language_id"`
	LocaleKeyGroup     string        `bson:"locale_key_group"`
	LocaleKey          string        `bson:"locale_key"`
	LocaleValue        string        `bson:"locale_value"`
}

type LocalizedPropertyRepository interface {
	CreateMany(c context.Context, items []LocalizedProperty) error
	Create(c context.Context, localization_settings *LocalizedProperty) error
	Update(c context.Context, localization_settings *LocalizedProperty) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocalizedProperty, error)
	FetchByID(c context.Context, ID string) (LocalizedProperty, error)
}

type LocalizedPropertyUsecase interface {
	CreateMany(c context.Context, items []LocalizedProperty) error
	FetchByID(c context.Context, ID string) (LocalizedProperty, error)
	Create(c context.Context, localization_settings *LocalizedProperty) error
	Update(c context.Context, localization_settings *LocalizedProperty) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocalizedProperty, error)
}
