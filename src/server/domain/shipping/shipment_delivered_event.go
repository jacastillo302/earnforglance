package domain

// ShipmentDeliveredEvent represents a shipment delivered event
type ShipmentDeliveredEvent struct {
	Shipment Shipment
}

// NewShipmentDeliveredEvent creates a new instance of ShipmentDeliveredEvent
func NewShipmentDeliveredEvent(shipment Shipment) *ShipmentDeliveredEvent {
	return &ShipmentDeliveredEvent{
		Shipment: shipment,
	}
}
