package domain

// SpecificationAttributeType represents a specification attribute type
type SpecificationAttributeType int

const (
	// Option represents an option type
	Option SpecificationAttributeType = 0

	// CustomText represents a custom text type
	CustomText SpecificationAttributeType = 10

	// CustomHtmlText represents a custom HTML text type
	CustomHtmlText SpecificationAttributeType = 20

	// Hyperlink represents a hyperlink type
	Hyperlink SpecificationAttributeType = 30
)
