package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGdprLog = "gdpr_logs"
)

// GdprLog represents a GDPR log
type GdprLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID     int                `bson:"customer_id"`
	ConsentID      int                `bson:"consent_id"`
	CustomerInfo   string             `bson:"customer_info"`
	RequestTypeID  int                `bson:"request_type_id"`
	RequestDetails string             `bson:"request_details"`
	CreatedOnUtc   time.Time          `bson:"created_on_utc"`
	RequestType    GdprRequestType    `bson:"request_type"`
}
