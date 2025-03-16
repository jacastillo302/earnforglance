package domain

// AddressField represents address fields
type AddressField int

const (
	// Country represents the country field
	Country AddressField = 0

	// StateProvince represents the state/province field
	StateProvince AddressField = 1

	// City represents the city field
	City AddressField = 2

	// County represents the county field
	County AddressField = 3

	// Address1 represents the address line 1 field
	Address1 AddressField = 4

	// Address2 represents the address line 2 field
	Address2 AddressField = 5

	// ZipPostalCode represents the zip/postal code field
	ZipPostalCode AddressField = 6
)
