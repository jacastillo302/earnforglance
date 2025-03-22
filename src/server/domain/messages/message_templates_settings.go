package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMessageTemplatesSettings = "message_templates_settings"
)

// MessageTemplatesSettings represents messages templates settings
type MessageTemplatesSettings struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	CaseInvariantReplacement bool               `bson:"case_invariant_replacement"`
	Color1                   string             `bson:"color1"`
	Color2                   string             `bson:"color2"`
	Color3                   string             `bson:"color3"`
}

// MessageTemplatesSettingsRepository interface
type MessageTemplatesSettingsRepository interface {
	Create(c context.Context, message_template_settings *MessageTemplatesSettings) error
	Update(c context.Context, message_template_settings *MessageTemplatesSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplatesSettings, error)
	FetchByID(c context.Context, ID string) (MessageTemplatesSettings, error)
}

// MessageTemplatesSettingsUsecase interface
type MessageTemplatesSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (MessageTemplatesSettings, error)
	Create(c context.Context, message_template_settings *MessageTemplatesSettings) error
	Update(c context.Context, message_template_settings *MessageTemplatesSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplatesSettings, error)
}
