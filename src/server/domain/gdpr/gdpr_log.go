package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionGdprLog = "gdpr_logs"
)

// GdprLog represents a GDPR logs
type GdprLog struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	CustomerID     bson.ObjectID `bson:"customer_id"`
	ConsentID      bson.ObjectID `bson:"consent_id"`
	CustomerInfo   string        `bson:"customer_info"`
	RequestTypeID  int           `bson:"request_type_id"`
	RequestDetails string        `bson:"request_details"`
	CreatedOnUtc   time.Time     `bson:"created_on_utc"`
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
	CreateMany(c context.Context, items []GdprLog) error
	FetchByID(c context.Context, ID string) (GdprLog, error)
	Create(c context.Context, gdpr_log *GdprLog) error
	Update(c context.Context, gdpr_log *GdprLog) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]GdprLog, error)
}
