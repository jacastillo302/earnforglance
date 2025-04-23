package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionShipmentItem = "shipment_items"
)

// ShipmentItem represents a shipment item
type ShipmentItem struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	ShipmentID  bson.ObjectID `bson:"shipment_id"`
	OrderItemID bson.ObjectID `bson:"order_item_id"`
	Quantity    int           `bson:"quantity"`
	WarehouseID bson.ObjectID `bson:"warehouse_id"`
}

// ShipmentItemRepository defines the repository interface for ShipmentItem
type ShipmentItemRepository interface {
	CreateMany(c context.Context, items []ShipmentItem) error
	Create(c context.Context, shipment_item *ShipmentItem) error
	Update(c context.Context, shipment_item *ShipmentItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShipmentItem, error)
	FetchByID(c context.Context, ID string) (ShipmentItem, error)
}

// ShipmentItemUsecase defines the use case interface for ShipmentItem
type ShipmentItemUsecase interface {
	CreateMany(c context.Context, items []ShipmentItem) error
	FetchByID(c context.Context, ID string) (ShipmentItem, error)
	Create(c context.Context, shipment_item *ShipmentItem) error
	Update(c context.Context, shipment_item *ShipmentItem) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ShipmentItem, error)
}
