package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionScheduleTask = "schedule_tasks"
)

// ScheduleTask represents a schedule task
type ScheduleTask struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Seconds        int                `bson:"seconds"`
	Type           string             `bson:"type"`
	LastEnabledUtc *time.Time         `bson:"last_enabled_utc,omitempty"`
	Enabled        bool               `bson:"enabled"`
	StopOnError    bool               `bson:"stop_on_error"`
	LastStartUtc   *time.Time         `bson:"last_start_utc,omitempty"`
	LastEndUtc     *time.Time         `bson:"last_end_utc,omitempty"`
	LastSuccessUtc *time.Time         `bson:"last_success_utc,omitempty"`
}
