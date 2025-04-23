package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionCustomerAttributeValue = "customer_attribute_values"
)

// CustomerAttributeValue represents a customer attribute value
type CustomerAttributeValue struct {
	ID                  bson.ObjectID `bson:"_id,omitempty"`
	CustomerAttributeID bson.ObjectID `bson:"customer_attribute_id"`
	Name                string        `bson:"name"`
	IsPreSelected       bool          `bson:"is_pre_selected"`
	DisplayOrder        int           `bson:"display_order"`
}

type CustomerAttributeValueRepository interface {
	CreateMany(c context.Context, items []CustomerAttributeValue) error
	Create(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Update(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAttributeValue, error)
	FetchByID(c context.Context, ID string) (CustomerAttributeValue, error)
}

type CustomerAttributeValueUsecase interface {
	CreateMany(c context.Context, items []CustomerAttributeValue) error
	FetchByID(c context.Context, ID string) (CustomerAttributeValue, error)
	Create(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Update(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]CustomerAttributeValue, error)
}
