package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionQueuedEmail = "queued_emails"
)

// QueuedEmail represents an email item
type QueuedEmail struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty"`
	PriorityID            int                 `bson:"priority_id"`
	From                  string              `bson:"from"`
	FromName              string              `bson:"from_name"`
	To                    string              `bson:"to"`
	ToName                string              `bson:"to_name"`
	ReplyTo               string              `bson:"reply_to"`
	ReplyToName           string              `bson:"reply_to_name"`
	CC                    string              `bson:"cc"`
	Bcc                   string              `bson:"bcc"`
	Subject               string              `bson:"subject"`
	Body                  string              `bson:"body"`
	AttachmentFilePath    string              `bson:"attachment_file_path"`
	AttachmentFileName    string              `bson:"attachment_file_name"`
	AttachedDownloadID    primitive.ObjectID  `bson:"attached_download_id"`
	CreatedOnUtc          time.Time           `bson:"created_on_utc"`
	DontSendBeforeDateUtc *time.Time          `bson:"dont_send_before_date_utc,omitempty"`
	SentTries             int                 `bson:"sent_tries"`
	SentOnUtc             *time.Time          `bson:"sent_on_utc,omitempty"`
	EmailAccountID        primitive.ObjectID  `bson:"email_account_id"`
	Priority              QueuedEmailPriority `bson:"priority"`
}

// QueuedEmailRepository interface
type QueuedEmailRepository interface {
	CreateMany(c context.Context, items []QueuedEmail) error
	Create(c context.Context, queued_email *QueuedEmail) error
	Update(c context.Context, queued_email *QueuedEmail) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]QueuedEmail, error)
	FetchByID(c context.Context, ID string) (QueuedEmail, error)
}

// QueuedEmailUsecase interface
type QueuedEmailUsecase interface {
	FetchByID(c context.Context, ID string) (QueuedEmail, error)
	Create(c context.Context, queued_email *QueuedEmail) error
	Update(c context.Context, queued_email *QueuedEmail) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]QueuedEmail, error)
}
