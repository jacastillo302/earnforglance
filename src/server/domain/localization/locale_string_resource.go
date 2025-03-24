package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLocaleStringResource = "locale_string_resources"
)

// LocaleStringResource represents a locale string resource
type LocaleStringResource struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID    primitive.ObjectID `bson:"language_id"`
	ResourceName  string             `bson:"resource_name"`
	ResourceValue string             `bson:"resource_value"`
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
	FetchByID(c context.Context, ID string) (LocaleStringResource, error)
	Create(c context.Context, locale_string_resource *LocaleStringResource) error
	Update(c context.Context, locale_string_resource *LocaleStringResource) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]LocaleStringResource, error)
}
