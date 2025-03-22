package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMessageTemplate = "message_templates"
)

// MessageTemplate represents a message template
type MessageTemplate struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	BccEmailAddresses  string             `bson:"bcc_email_addresses"`
	Subject            string             `bson:"subject"`
	Body               string             `bson:"body"`
	IsActive           bool               `bson:"is_active"`
	DelayBeforeSend    *int               `bson:"delay_before_send,omitempty"`
	DelayPeriodID      int                `bson:"delay_period_id"`
	AttachedDownloadID primitive.ObjectID `bson:"attached_download_id"`
	AllowDirectReply   bool               `bson:"allow_direct_reply"`
	EmailAccountID     primitive.ObjectID `bson:"email_account_id"`
	LimitedToStores    bool               `bson:"limited_to_stores"`
	DelayPeriod        MessageDelayPeriod `bson:"delay_period"`
}

// MessageTemplateRepository represents the repository interface for MessageTemplate
type MessageTemplateRepository interface {
	Create(c context.Context, message_template *MessageTemplate) error
	Update(c context.Context, message_template *MessageTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplate, error)
	FetchByID(c context.Context, ID string) (MessageTemplate, error)
}

// MessageTemplateUsecase represents the usecase interface for MessageTemplate
type MessageTemplateUsecase interface {
	FetchByID(c context.Context, ID string) (MessageTemplate, error)
	Create(c context.Context, message_template *MessageTemplate) error
	Update(c context.Context, message_template *MessageTemplate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessageTemplate, error)
}
