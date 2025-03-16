package domain

// ShipmentCreatedEvent represents a shipment created event
type ShipmentCreatedEvent struct {
	Shipment Shipment
}

// NewShipmentCreatedEvent creates a new instance of ShipmentCreatedEvent
func NewShipmentCreatedEvent(shipment Shipment) *ShipmentCreatedEvent {
	return &ShipmentCreatedEvent{
		Shipment: shipment,
	}
}
