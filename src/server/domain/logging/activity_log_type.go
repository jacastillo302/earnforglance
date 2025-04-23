package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionActivityLogType = "activity_log_types"
)

// ActivityLogType represents an activity log type records
type ActivityLogType struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	SystemKeyword string        `bson:"system_keyword"`
	Name          string        `bson:"name"`
	Enabled       bool          `bson:"enabled"`
}

type ActivityLogTypeRepository interface {
	CreateMany(c context.Context, items []ActivityLogType) error
	Create(c context.Context, activity_log_type *ActivityLogType) error
	Update(c context.Context, activity_log_type *ActivityLogType) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ActivityLogType, error)
	FetchByID(c context.Context, ID string) (ActivityLogType, error)
}

type ActivityLogTypeUsecase interface {
	CreateMany(c context.Context, items []ActivityLogType) error
	FetchByID(c context.Context, ID string) (ActivityLogType, error)
	Create(c context.Context, activity_log_type *ActivityLogType) error
	Update(c context.Context, activity_log_type *ActivityLogType) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ActivityLogType, error)
}
