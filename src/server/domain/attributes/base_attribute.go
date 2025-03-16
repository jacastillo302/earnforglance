package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// AttributeControlType represents the type of attribute control
type AttributeControlType int

const (
	TextBox AttributeControlType = iota
	MultilineTextbox
	Datepicker
	FileUpload
	ReadonlyCheckboxes
)

const (
	CollectionBaseAttribute = "base_attributes"
)

// BaseAttribute represents the base class for attributes
type BaseAttribute struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name"`
	IsRequired             bool               `bson:"is_required"`
	AttributeControlTypeId int                `bson:"attribute_control_type_id"`
	DisplayOrder           int                `bson:"display_order"`
}

// AttributeControlType returns the attribute control type
func (ba *BaseAttribute) AttributeControlType() AttributeControlType {
	return AttributeControlType(ba.AttributeControlTypeId)
}

// SetAttributeControlType sets the attribute control type
func (ba *BaseAttribute) SetAttributeControlType(value AttributeControlType) {
	ba.AttributeControlTypeId = int(value)
}

// ShouldHaveValues indicates whether this attribute should have values
func (ba *BaseAttribute) ShouldHaveValues() bool {
	switch ba.AttributeControlType() {
	case TextBox, MultilineTextbox, Datepicker, FileUpload:
		return false
	}
	return true
}

// CanBeUsedAsCondition indicates whether this attribute can be used as condition for some other attribute
func (ba *BaseAttribute) CanBeUsedAsCondition() bool {
	switch ba.AttributeControlType() {
	case ReadonlyCheckboxes, TextBox, MultilineTextbox, Datepicker, FileUpload:
		return false
	}
	return true
}
