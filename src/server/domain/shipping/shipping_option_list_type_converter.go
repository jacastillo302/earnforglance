package domain

import (
	"bytes"
	"encoding/xml"
)

// ShippingOptionListTypeConverter is a type converter for a list of ShippingOption
type ShippingOptionListTypeConverter struct{}

// CanConvertFrom checks if the source type can be converted to a list of ShippingOption
func (c *ShippingOptionListTypeConverter) CanConvertFrom(sourceType string) bool {
	return sourceType == "string"
}

// ConvertFrom converts the given value to a list of ShippingOption
func (c *ShippingOptionListTypeConverter) ConvertFrom(value string) ([]ShippingOption, error) {
	if value == "" {
		return nil, nil
	}

	var shippingOptions []ShippingOption
	err := xml.Unmarshal([]byte(value), &shippingOptions)
	if err != nil {
		return nil, err
	}

	return shippingOptions, nil
}

// ConvertTo converts the given list of ShippingOption to a string
func (c *ShippingOptionListTypeConverter) ConvertTo(shippingOptions []ShippingOption) (string, error) {
	if shippingOptions == nil {
		return "", nil
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(shippingOptions)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
