package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPermisionRecordAttributeValue = "permission_record_attribute_values"
)

// PermisionRecordAttributeValue represents a customer attribute value
type PermisionRecordAttributeValue struct {
	ID                         bson.ObjectID `bson:"_id,omitempty"`
	PermisionRecordAttributeID bson.ObjectID `bson:"permission_record_attribute_id"`
	RecordID                   bson.ObjectID `bson:"record_id"`
	Value                      string        `bson:"value"`
	IsPreSelected              bool          `bson:"is_pre_selected"`
	DisplayOrder               int           `bson:"display_order"`
}

type PermisionRecordAttributeValueRepository interface {
	CreateMany(c context.Context, items []PermisionRecordAttributeValue) error
	Create(c context.Context, customer_attribute_value *PermisionRecordAttributeValue) error
	Update(c context.Context, customer_attribute_value *PermisionRecordAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermisionRecordAttributeValue, error)
	FetchByID(c context.Context, ID string) (PermisionRecordAttributeValue, error)
}

type PermisionRecordAttributeValueUsecase interface {
	CreateMany(c context.Context, items []PermisionRecordAttributeValue) error
	FetchByID(c context.Context, ID string) (PermisionRecordAttributeValue, error)
	Create(c context.Context, customer_attribute_value *PermisionRecordAttributeValue) error
	Update(c context.Context, customer_attribute_value *PermisionRecordAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PermisionRecordAttributeValue, error)
}
