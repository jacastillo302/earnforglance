package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionGdprLog = "gdpr_logs"
)

// GdprLog represents a GDPR logs
type GdprLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID     primitive.ObjectID `bson:"customer_id"`
	ConsentID      primitive.ObjectID `bson:"consent_id"`
	CustomerInfo   string             `bson:"customer_info"`
	RequestTypeID  int                `bson:"request_type_id"`
	RequestDetails string             `bson:"request_details"`
	CreatedOnUtc   time.Time          `bson:"created_on_utc"`
	RequestType    GdprRequestType    `bson:"request_type"`
}

// GdprLogRepository interface
type GdprLogRepository interface {
	CreateMany(c context.Context, items []GdprLog) error
	Create(c context.Context, gdpr_log *GdprLog) error
	Update(c context.Context, gdpr_log *GdprLog) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprLog, error)
	FetchByID(c context.Context, ID string) (GdprLog, error)
}

// GdprLogUsecase interface
type GdprLogUsecase interface {
	FetchByID(c context.Context, ID string) (GdprLog, error)
	Create(c context.Context, gdpr_log *GdprLog) error
	Update(c context.Context, gdpr_log *GdprLog) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprLog, error)
}
