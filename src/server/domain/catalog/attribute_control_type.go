package domain

import "encoding/json"

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

func (a AttributeControlType) String() string {
	switch a {
	case DropdownList:
		return "DropdownList"
	case RadioList:
		return "RadioList"
	case Checkboxes:
		return "Checkboxes"
	case TextBox:
		return "TextBox"
	case MultilineTextbox:
		return "MultilineTextbox"
	case Datepicker:
		return "Datepicker"
	case FileUpload:
		return "FileUpload"
	case ColorSquares:
		return "ColorSquares"
	case ImageSquares:
		return "ImageSquares"
	case ReadonlyCheckboxes:
		return "ReadonlyCheckboxes"
	default:
		return "Unknown"
	}
}

func (a AttributeControlType) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}
