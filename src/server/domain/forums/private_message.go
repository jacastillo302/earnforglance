package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPrivateMessage = "private_messages"
)

// PrivateMessage represents a private message
type PrivateMessage struct {
	ID                   bson.ObjectID `bson:"_id,omitempty"`
	StoreID              bson.ObjectID `bson:"store_id"`
	FromCustomerID       bson.ObjectID `bson:"from_customer_id"`
	ToCustomerID         bson.ObjectID `bson:"to_customer_id"`
	Subject              string        `bson:"subject"`
	Text                 string        `bson:"text"`
	IsRead               bool          `bson:"is_read"`
	IsDeletedByAuthor    bool          `bson:"is_deleted_by_author"`
	IsDeletedByRecipient bool          `bson:"is_deleted_by_recipient"`
	CreatedOnUtc         time.Time     `bson:"created_on_utc"`
}

// PrivateMessageRepository interface
type PrivateMessageRepository interface {
	CreateMany(c context.Context, items []PrivateMessage) error
	Create(c context.Context, private_message *PrivateMessage) error
	Update(c context.Context, private_message *PrivateMessage) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PrivateMessage, error)
	FetchByID(c context.Context, ID string) (PrivateMessage, error)
}

// PrivateMessageUsecase interface
type PrivateMessageUsecase interface {
	CreateMany(c context.Context, items []PrivateMessage) error
	FetchByID(c context.Context, ID string) (PrivateMessage, error)
	Create(c context.Context, private_message *PrivateMessage) error
	Update(c context.Context, private_message *PrivateMessage) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PrivateMessage, error)
}
