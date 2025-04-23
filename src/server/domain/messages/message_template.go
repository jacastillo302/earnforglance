package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionMessageTemplate = "message_templates"
)

// MessageTemplate represents a message template
type MessageTemplate struct {
	ID                 bson.ObjectID `bson:"_id,omitempty"`
	Name               string        `bson:"name"`
	BccEmailAddresses  string        `bson:"bcc_email_addresses"`
	Subject            string        `bson:"subject"`
	Body               string        `bson:"body"`
	IsActive           bool          `bson:"is_active"`
	DelayBeforeSend    *int          `bson:"delay_before_send"`
	DelayPeriodID      int           `bson:"delay_period_id"`
	AttachedDownloadID string        `bson:"attached_download_id"`
	AllowDirectReply   bool          `bson:"allow_direct_reply"`
	EmailAccountID     bson.ObjectID `bson:"email_account_id"`
	LimitedToStores    bool          `bson:"limited_to_stores"`
}

// MessageTemplateRepository represents the repository interface for MessageTemplate
type MessageTemplateRepository interface {
	CreateMany(c context.Context, items []MessageTemplate) error
	Create(c context.Context, message_template *MessageTemplate) error
	Update(c context.Context, message_template *MessageTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplate, error)
	FetchByID(c context.Context, ID string) (MessageTemplate, error)
}

// MessageTemplateUsecase represents the usecase interface for MessageTemplate
type MessageTemplateUsecase interface {
	CreateMany(c context.Context, items []MessageTemplate) error
	FetchByID(c context.Context, ID string) (MessageTemplate, error)
	Create(c context.Context, message_template *MessageTemplate) error
	Update(c context.Context, message_template *MessageTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplate, error)
}
