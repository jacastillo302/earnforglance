package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionLocaleStringResource = "locale_string_resources"
)

// LocaleStringResource represents a locale string resource
type LocaleStringResource struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	LanguageID    bson.ObjectID `bson:"language_id"`
	ResourceName  string        `bson:"resource_name"`
	ResourceValue string        `bson:"resource_value"`
}

type LocaleStringResourceRepository interface {
	CreateMany(c context.Context, items []LocaleStringResource) error
	Create(c context.Context, locale_string_resource *LocaleStringResource) error
	Update(c context.Context, locale_string_resource *LocaleStringResource) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocaleStringResource, error)
	FetchByID(c context.Context, ID string) (LocaleStringResource, error)
}

type LocaleStringResourceUsecase interface {
	CreateMany(c context.Context, items []LocaleStringResource) error
	FetchByID(c context.Context, ID string) (LocaleStringResource, error)
	Create(c context.Context, locale_string_resource *LocaleStringResource) error
	Update(c context.Context, locale_string_resource *LocaleStringResource) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocaleStringResource, error)
}
