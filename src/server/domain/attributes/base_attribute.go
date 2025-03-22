package domain

import (
	"context"

	catalog "earnforglance/server/domain/catalog"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type BaseAttributeRepository interface {
	Create(c context.Context, base_attribute *BaseAttribute) error
	Update(c context.Context, base_attribute *BaseAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BaseAttribute, error)
	FetchByID(c context.Context, ID string) (BaseAttribute, error)
}

type BaseAttributeUsecase interface {
	FetchByID(c context.Context, ID string) (BaseAttribute, error)
	Create(c context.Context, base_attribute *BaseAttribute) error
	Update(c context.Context, base_attribute *BaseAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]BaseAttribute, error)
}

// AttributeControlType returns the attribute control type
func (ba *BaseAttribute) AttributeControlType() catalog.AttributeControlType {
	return catalog.AttributeControlType(ba.AttributeControlTypeId)
}

// SetAttributeControlType sets the attribute control type
func (ba *BaseAttribute) SetAttributeControlType(value catalog.AttributeControlType) {
	ba.AttributeControlTypeId = int(value)
}

// ShouldHaveValues indicates whether this attribute should have values
func (ba *BaseAttribute) ShouldHaveValues() bool {
	switch ba.AttributeControlType() {
	case catalog.TextBox, catalog.MultilineTextbox, catalog.Datepicker, catalog.FileUpload:
		return false
	}
	return true
}

// CanBeUsedAsCondition indicates whether this attribute can be used as condition for some other attribute
func (ba *BaseAttribute) CanBeUsedAsCondition() bool {
	switch ba.AttributeControlType() {
	case catalog.ReadonlyCheckboxes, catalog.TextBox, catalog.MultilineTextbox, catalog.Datepicker, catalog.FileUpload:
		return false
	}
	return true
}
