package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionLog = "logs"
)

// Log represents a log record
type Log struct {
	ID           bson.ObjectID  `bson:"_id,omitempty"`
	LogLevelID   int            `bson:"log_level_id"`
	ShortMessage string         `bson:"short_message"`
	FullMessage  string         `bson:"full_message"`
	IpAddress    string         `bson:"ip_address"`
	CustomerID   *bson.ObjectID `bson:"customer_id"`
	PageUrl      string         `bson:"page_url"`
	ReferrerUrl  string         `bson:"referrer_url"`
	CreatedOnUtc time.Time      `bson:"created_on_utc"`
}

type LogRepository interface {
	CreateMany(c context.Context, items []Log) error
	Create(c context.Context, log *Log) error
	Update(c context.Context, log *Log) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Log, error)
	FetchByID(c context.Context, ID string) (Log, error)
}

type LogUsecase interface {
	CreateMany(c context.Context, items []Log) error
	FetchByID(c context.Context, ID string) (Log, error)
	Create(c context.Context, log *Log) error
	Update(c context.Context, log *Log) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Log, error)
}
