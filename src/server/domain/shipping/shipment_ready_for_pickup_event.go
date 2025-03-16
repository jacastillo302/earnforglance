package domain

// ShipmentReadyForPickupEvent represents a shipment ready for pickup event
type ShipmentReadyForPickupEvent struct {
	Shipment Shipment
}

// NewShipmentReadyForPickupEvent creates a new instance of ShipmentReadyForPickupEvent
func NewShipmentReadyForPickupEvent(shipment Shipment) *ShipmentReadyForPickupEvent {
	return &ShipmentReadyForPickupEvent{
		Shipment: shipment,
	}
}
