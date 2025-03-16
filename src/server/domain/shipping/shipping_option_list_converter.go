package domain

import (
	"bytes"
	"encoding/xml"
)

// ShippingOptionTypeConverter is a type converter for ShippingOption
type ShippingOptionTypeConverter struct{}

// CanConvertFrom checks if the source type can be converted to ShippingOption
func (c *ShippingOptionTypeConverter) CanConvertFrom(sourceType string) bool {
	return sourceType == "string"
}

// ConvertFrom converts the given value to ShippingOption
func (c *ShippingOptionTypeConverter) ConvertFrom(value string) (*ShippingOption, error) {
	if value == "" {
		return nil, nil
	}

	var shippingOption ShippingOption
	err := xml.Unmarshal([]byte(value), &shippingOption)
	if err != nil {
		return nil, err
	}

	return &shippingOption, nil
}

// ConvertTo converts the given ShippingOption to a string
func (c *ShippingOptionTypeConverter) ConvertTo(shippingOption *ShippingOption) (string, error) {
	if shippingOption == nil {
		return "", nil
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(shippingOption)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
