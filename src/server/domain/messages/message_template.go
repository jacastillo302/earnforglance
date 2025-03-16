package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	AttachedDownloadID int                `bson:"attached_download_id"`
	AllowDirectReply   bool               `bson:"allow_direct_reply"`
	EmailAccountID     int                `bson:"email_account_id"`
	LimitedToStores    bool               `bson:"limited_to_stores"`
	DelayPeriod        MessageDelayPeriod `bson:"delay_period"`
}
