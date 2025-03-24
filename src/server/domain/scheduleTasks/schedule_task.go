package domain

import (
	"context" // Added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionScheduleTask = "schedule_tasks"
)

// ScheduleTask represents a schedule task.
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

// ScheduleTaskRepository defines the repository interface for ScheduleTask
type ScheduleTaskRepository interface {
	CreateMany(c context.Context, items []ScheduleTask) error
	Create(c context.Context, schedule_task *ScheduleTask) error
	Update(c context.Context, schedule_task *ScheduleTask) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ScheduleTask, error)
	FetchByID(c context.Context, ID string) (ScheduleTask, error)
}

// ScheduleTaskUsecase defines the usecase interface for ScheduleTask
type ScheduleTaskUsecase interface {
	CreateMany(c context.Context, items []ScheduleTask) error
	FetchByID(c context.Context, ID string) (ScheduleTask, error)
	Create(c context.Context, schedule_task *ScheduleTask) error
	Update(c context.Context, schedule_task *ScheduleTask) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ScheduleTask, error)
}
