package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionWidgetSettings = "widget_settings"
)

// WidgetSettings represents widget settings
type WidgetSettings struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	ActiveWidgetSystemNames []string           `bson:"active_widget_system_names"`
}

// NewWidgetSettings creates a new WidgetSettings instance
func NewWidgetSettings() *WidgetSettings {
	return &WidgetSettings{
		ActiveWidgetSystemNames: []string{},
	}
}
