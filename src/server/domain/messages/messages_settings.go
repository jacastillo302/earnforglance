package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionMessagesSettings = "messages_settings"
)

// MessagesSettings represents messages settings
type MessagesSettings struct {
	ID                                            primitive.ObjectID `bson:"_id,omitempty"`
	UsePopupNotifications                         bool               `bson:"use_popup_notifications"`
	UseDefaultEmailAccountForSendStoreOwnerEmails bool               `bson:"use_default_email_account_for_send_store_owner_emails"`
}
