package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPermisionRecordAttribute = "permission_record_attributes"
)

// PermisionRecordAttribute represents a customer attribute
type PermisionRecordAttribute struct {
	ID                              bson.ObjectID `bson:"_id,omitempty"`
	BaseAttributeID                 bson.ObjectID `bson:"base_attribute_id,omitempty"`
	PermissionRecordID              bson.ObjectID `bson:"permission_record_id,omitempty"`
	Name                            string        `bson:"name"`
	IsRequired                      bool          `bson:"is_required"`
	DisplayOrder                    int           `bson:"display_order"`
	DefaultValue                    string        `bson:"default_value"`
	ValidationMinLength             *int          `bson:"validation_min_length"`
	ValidationMaxLength             *int          `bson:"validation_max_length"`
	ValidationFileAllowedExtensions string        `bson:"validation_file_allowed_extensions"`
	ValidationFileMaximumSize       *int          `bson:"validation_file_maximum_size"`
	ConditionAttributeXml           string        `bson:"condition_attribute_xml"`
}

type PermisionRecordAttributeRepository interface {
	CreateMany(c context.Context, items []PermisionRecordAttribute) error
	Create(c context.Context, customer_attribute *PermisionRecordAttribute) error
	Update(c context.Context, customer_attribute *PermisionRecordAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermisionRecordAttribute, error)
	FetchByID(c context.Context, ID string) (PermisionRecordAttribute, error)
}

type PermisionRecordAttributeUsecase interface {
	CreateMany(c context.Context, items []PermisionRecordAttribute) error
	FetchByID(c context.Context, ID string) (PermisionRecordAttribute, error)
	Create(c context.Context, customer_attribute *PermisionRecordAttribute) error
	Update(c context.Context, customer_attribute *PermisionRecordAttribute) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermisionRecordAttribute, error)
}
