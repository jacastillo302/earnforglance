package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCampaigno = "Campaigns"
)

// Campaign represents a Campaign
type Campaign struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Name                  string             `bson:"name"`
	Subject               string             `bson:"subject"`
	Body                  string             `bson:"body"`
	StoreID               int                `bson:"store_id"`
	CustomerRoleID        int                `bson:"customer_role_id"`
	CreatedOnUtc          time.Time          `bson:"created_on_utc"`
	DontSendBeforeDateUtc *time.Time         `bson:"dont_send_before_date_utc,omitempty"`
}
