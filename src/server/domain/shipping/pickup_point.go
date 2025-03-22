package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPickupPoint = "pickup_points"
)

// PickupPoint represents a pickup point.
type PickupPoint struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	PickupPointID      string             `bson:"pickup_point_id"`
	Name               string             `bson:"name"`
	Description        string             `bson:"description"`
	ProviderSystemName string             `bson:"provider_system_name"`
	Address            string             `bson:"address"`
	City               string             `bson:"city"`
	County             string             `bson:"county"`
	StateAbbreviation  string             `bson:"state_abbreviation"`
	CountryCode        string             `bson:"country_code"`
	ZipPostalCode      string             `bson:"zip_postal_code"`
	Latitude           *float64           `bson:"latitude,omitempty"`
	Longitude          *float64           `bson:"longitude,omitempty"`
	PickupFee          float64            `bson:"pickup_fee"`
	OpeningHours       string             `bson:"opening_hours"`
	DisplayOrder       int                `bson:"display_order"`
	TransitDays        *int               `bson:"transit_days,omitempty"`
}

// PickupPointRepository defines the repository interface for PickupPoint
type PickupPointRepository interface {
	Create(c context.Context, pickup_point *PickupPoint) error
	Update(c context.Context, pickup_point *PickupPoint) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PickupPoint, error)
	FetchByID(c context.Context, ID string) (PickupPoint, error)
}

// PickupPointUsecase defines the use case interface for PickupPoint
type PickupPointUsecase interface {
	FetchByID(c context.Context, ID string) (PickupPoint, error)
	Create(c context.Context, pickup_point *PickupPoint) error
	Update(c context.Context, pickup_point *PickupPoint) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PickupPoint, error)
}
