package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionActivityLogType = "activity_log_types"
)

// ActivityLogType represents an activity log type record
type ActivityLogType struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	SystemKeyword string             `bson:"system_keyword"`
	Name          string             `bson:"name"`
	Enabled       bool               `bson:"enabled"`
}
