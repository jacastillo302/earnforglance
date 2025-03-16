package domain

// AttributeControlType represents an attribute control type
type AttributeControlType int

const (
	// DropdownList represents a dropdown list
	DropdownList AttributeControlType = 1

	// RadioList represents a radio list
	RadioList AttributeControlType = 2

	// Checkboxes represents checkboxes
	Checkboxes AttributeControlType = 3

	// TextBox represents a text box
	TextBox AttributeControlType = 4

	// MultilineTextbox represents a multiline text box
	MultilineTextbox AttributeControlType = 10

	// Datepicker represents a date picker
	Datepicker AttributeControlType = 20

	// FileUpload represents a file upload control
	FileUpload AttributeControlType = 30

	// ColorSquares represents color squares
	ColorSquares AttributeControlType = 40

	// ImageSquares represents image squares
	ImageSquares AttributeControlType = 45

	// ReadonlyCheckboxes represents read-only checkboxes
	ReadonlyCheckboxes AttributeControlType = 50
)
