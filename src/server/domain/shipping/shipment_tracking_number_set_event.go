package domain

// ShipmentTrackingNumberSetEvent represents a shipment tracking number set event
type ShipmentTrackingNumberSetEvent struct {
	Shipment Shipment
}

// NewShipmentTrackingNumberSetEvent creates a new instance of ShipmentTrackingNumberSetEvent
func NewShipmentTrackingNumberSetEvent(shipment Shipment) *ShipmentTrackingNumberSetEvent {
	return &ShipmentTrackingNumberSetEvent{
		Shipment: shipment,
	}
}
