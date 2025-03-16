package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionActivityLog = "activity_logs"
)

// ActivityLog represents an activity log record
type ActivityLog struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ActivityLogTypeID int                `bson:"activity_log_type_id"`
	EntityID          *int               `bson:"entity_id,omitempty"`
	EntityName        string             `bson:"entity_name"`
	CustomerID        int                `bson:"customer_id"`
	Comment           string             `bson:"comment"`
	CreatedOnUtc      time.Time          `bson:"created_on_utc"`
	IpAddress         string             `bson:"ip_address"`
}
