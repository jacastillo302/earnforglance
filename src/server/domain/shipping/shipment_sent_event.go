package domain

// ShipmentSentEvent represents a shipment sent event
type ShipmentSentEvent struct {
	Shipment Shipment
}

// NewShipmentSentEvent creates a new instance of ShipmentSentEvent
func NewShipmentSentEvent(shipment Shipment) *ShipmentSentEvent {
	return &ShipmentSentEvent{
		Shipment: shipment,
	}
}
