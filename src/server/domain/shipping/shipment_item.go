package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionShipmentItem = "shipment_items"
)

// ShipmentItem represents a shipment item
type ShipmentItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ShipmentID  int                `bson:"shipment_id"`
	OrderItemID int                `bson:"order_item_id"`
	Quantity    int                `bson:"quantity"`
	WarehouseID int                `bson:"warehouse_id"`
}
