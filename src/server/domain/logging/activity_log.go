package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionActivityLog = "activity_logs"
)

// ActivityLog represents an activity log records
type ActivityLog struct {
	ID                bson.ObjectID  `bson:"_id,omitempty"`
	ActivityLogTypeID bson.ObjectID  `bson:"activity_log_type_id"`
	EntityID          *bson.ObjectID `bson:"entity_id"`
	EntityName        string         `bson:"entity_name"`
	CustomerID        bson.ObjectID  `bson:"customer_id"`
	Comment           string         `bson:"comment"`
	CreatedOnUtc      time.Time      `bson:"created_on_utc"`
	IpAddress         string         `bson:"ip_address"`
}

type ActivityLogRepository interface {
	CreateMany(c context.Context, items []ActivityLog) error
	Create(c context.Context, activity_log *ActivityLog) error
	Update(c context.Context, activity_log *ActivityLog) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ActivityLog, error)
	FetchByID(c context.Context, ID string) (ActivityLog, error)
}

type ActivityLogUsecase interface {
	CreateMany(c context.Context, items []ActivityLog) error
	FetchByID(c context.Context, ID string) (ActivityLog, error)
	Create(c context.Context, activity_log *ActivityLog) error
	Update(c context.Context, activity_log *ActivityLog) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ActivityLog, error)
}
