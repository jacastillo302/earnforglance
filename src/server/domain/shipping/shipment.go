package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionShipment = "shipments"
)

// Shipment represents a shipment
type Shipment struct {
	ID                    bson.ObjectID `bson:"_id,omitempty"`
	OrderID               bson.ObjectID `bson:"order_id"`
	TrackingNumber        string        `bson:"tracking_number"`
	TotalWeight           *float64      `bson:"total_weight"`
	ShippedDateUtc        *time.Time    `bson:"shipped_date_utc"`
	DeliveryDateUtc       *time.Time    `bson:"delivery_date_utc"`
	ReadyForPickupDateUtc *time.Time    `bson:"ready_for_pickup_date_utc"`
	AdminComment          string        `bson:"admin_comment"`
	CreatedOnUtc          time.Time     `bson:"created_on_utc"`
}

// ShipmentRepository defines the repository interface for Shipment
type ShipmentRepository interface {
	CreateMany(c context.Context, items []Shipment) error
	Create(c context.Context, shipment *Shipment) error
	Update(c context.Context, shipment *Shipment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Shipment, error)
	FetchByID(c context.Context, ID string) (Shipment, error)
}

// ShipmentUsecase defines the use case interface for Shipment
type ShipmentUsecase interface {
	CreateMany(c context.Context, items []Shipment) error
	FetchByID(c context.Context, ID string) (Shipment, error)
	Create(c context.Context, shipment *Shipment) error
	Update(c context.Context, shipment *Shipment) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Shipment, error)
}
