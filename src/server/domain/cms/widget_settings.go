package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionWidgetSettings = "widget_settings"
)

// WidgetSettings represents widget settings
type WidgetSettings struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	ActiveWidgetSystemNames []string           `bson:"active_widget_system_names"`
}

type WidgetSettingsRepository interface {
	Create(c context.Context, widget_settings *WidgetSettings) error
	Update(c context.Context, widget_settings *WidgetSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]WidgetSettings, error)
	FetchByID(c context.Context, ID string) (WidgetSettings, error)
}

type WidgetSettingsUsecase interface {
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
