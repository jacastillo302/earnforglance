package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPrivateMessage = "private_messages"
)

// PrivateMessage represents a private message
type PrivateMessage struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	StoreID              int                `bson:"store_id"`
	FromCustomerID       int                `bson:"from_customer_id"`
	ToCustomerID         int                `bson:"to_customer_id"`
	Subject              string             `bson:"subject"`
	Text                 string             `bson:"text"`
	IsRead               bool               `bson:"is_read"`
	IsDeletedByAuthor    bool               `bson:"is_deleted_by_author"`
	IsDeletedByRecipient bool               `bson:"is_deleted_by_recipient"`
	CreatedOnUtc         time.Time          `bson:"created_on_utc"`
}
