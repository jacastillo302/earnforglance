package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionWidgetSettings = "widget_settings"
)

// WidgetSettings represents widget settings
type WidgetSettings struct {
	ID                      bson.ObjectID `bson:"_id,omitempty"`
	ActiveWidgetSystemNames []string      `bson:"active_widget_system_names"`
}

type WidgetSettingsRepository interface {
	CreateMany(c context.Context, items []WidgetSettings) error
	Create(c context.Context, widget_settings *WidgetSettings) error
	Update(c context.Context, widget_settings *WidgetSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]WidgetSettings, error)
	FetchByID(c context.Context, ID string) (WidgetSettings, error)
}

type WidgetSettingsUsecase interface {
	CreateMany(c context.Context, items []WidgetSettings) error
	FetchByID(c context.Context, ID string) (WidgetSettings, error)
	Create(c context.Context, widget_settings *WidgetSettings) error
	Update(c context.Context, widget_settings *WidgetSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]WidgetSettings, error)
}

// NewWidgetSettings creates a new WidgetSettings instance
func NewWidgetSettings() *WidgetSettings {
	return &WidgetSettings{
		ActiveWidgetSystemNames: []string{},
	}
}
