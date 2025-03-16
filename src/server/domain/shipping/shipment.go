package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionShipment = "shipments"
)

// Shipment represents a shipment
type Shipment struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	OrderID               int                `bson:"order_id"`
	TrackingNumber        string             `bson:"tracking_number"`
	TotalWeight           *float64           `bson:"total_weight,omitempty"`
	ShippedDateUtc        *time.Time         `bson:"shipped_date_utc,omitempty"`
	DeliveryDateUtc       *time.Time         `bson:"delivery_date_utc,omitempty"`
	ReadyForPickupDateUtc *time.Time         `bson:"ready_for_pickup_date_utc,omitempty"`
	AdminComment          string             `bson:"admin_comment"`
	CreatedOnUtc          time.Time          `bson:"created_on_utc"`
}
