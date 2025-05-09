package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionMessagesSettings = "messages_settings"
)

// MessagesSettings represents message settings
type MessagesSettings struct {
	ID                                            bson.ObjectID `bson:"_id,omitempty"`
	UsePopupNotifications                         bool          `bson:"use_popup_notifications"`
	UseDefaultEmailAccountForSendStoreOwnerEmails bool          `bson:"use_default_email_account_for_send_store_owner_emails"`
}

type MessagesSettingsRepository interface {
	CreateMany(c context.Context, items []MessagesSettings) error
	Create(c context.Context, messages_settings *MessagesSettings) error
	Update(c context.Context, messages_settings *MessagesSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessagesSettings, error)
	FetchByID(c context.Context, ID string) (MessagesSettings, error)
}

type MessagesSettingsUsecase interface {
	CreateMany(c context.Context, items []MessagesSettings) error
	FetchByID(c context.Context, ID string) (MessagesSettings, error)
	Create(c context.Context, messages_settings *MessagesSettings) error
	Update(c context.Context, messages_settings *MessagesSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MessagesSettings, error)
}
