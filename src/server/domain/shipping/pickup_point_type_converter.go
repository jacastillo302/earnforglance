package domain

import (
	"encoding/xml"
)

// PickupPointTypeConverter is a type converter for PickupPoint
type PickupPointTypeConverter struct{}

// CanConvertFrom checks if the source type can be converted to PickupPoint
func (c *PickupPointTypeConverter) CanConvertFrom(sourceType string) bool {
	return sourceType == "string"
}

// ConvertFrom converts the given value to PickupPoint
func (c *PickupPointTypeConverter) ConvertFrom(value string) (*PickupPoint, error) {
	if value == "" {
		return nil, nil
	}

	var pickupPoint PickupPoint
	err := xml.Unmarshal([]byte(value), &pickupPoint)
	if err != nil {
		return nil, err
	}

	return &pickupPoint, nil
}

// ConvertTo converts the given PickupPoint to a string
func (c *PickupPointTypeConverter) ConvertTo(pickupPoint *PickupPoint) (string, error) {
	if pickupPoint == nil {
		return "", nil
	}

	output, err := xml.Marshal(pickupPoint)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
