package domain

import (
	"context" // added context library
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCampaign = "Campaigns"
)

// Campaign represents a Campaign
type Campaign struct {
	ID                    bson.ObjectID `bson:"_id,omitempty"`
	Name                  string        `bson:"name"`
	Subject               string        `bson:"subject"`
	Body                  string        `bson:"body"`
	StoreID               bson.ObjectID `bson:"store_id"`
	CustomerRoleID        bson.ObjectID `bson:"customer_role_id"`
	CreatedOnUtc          time.Time     `bson:"created_on_utc"`
	DontSendBeforeDateUtc *time.Time    `bson:"dont_send_before_date_utc"`
}

// CampaignRepository represents the repository interface for Campaign
type CampaignRepository interface {
	CreateMany(c context.Context, items []Campaign) error
	Create(c context.Context, campaign *Campaign) error
	Update(c context.Context, campaign *Campaign) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Campaign, error)
	FetchByID(c context.Context, ID string) (Campaign, error)
}

// CampaignUsecase represents the use case interface for Campaign
type CampaignUsecase interface {
	CreateMany(c context.Context, items []Campaign) error
	FetchByID(c context.Context, ID string) (Campaign, error)
	Create(c context.Context, campaign *Campaign) error
	Update(c context.Context, campaign *Campaign) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Campaign, error)
}
