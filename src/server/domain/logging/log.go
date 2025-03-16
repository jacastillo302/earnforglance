package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLog = "logs"
)

// Log represents a log record
type Log struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	LogLevelID   int                `bson:"log_level_id"`
	ShortMessage string             `bson:"short_message"`
	FullMessage  string             `bson:"full_message"`
	IpAddress    string             `bson:"ip_address"`
	CustomerID   *int               `bson:"customer_id,omitempty"`
	PageUrl      string             `bson:"page_url"`
	ReferrerUrl  string             `bson:"referrer_url"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
	LogLevel     LogLevel           `bson:"log_level"`
}
