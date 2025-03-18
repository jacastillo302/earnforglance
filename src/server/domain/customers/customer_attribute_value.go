package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCustomerAttributeValue = "customer_attribute_values"
)

// CustomerAttributeValue represents a customer attribute value
type CustomerAttributeValue struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	CustomerAttributeID primitive.ObjectID `bson:"customer_attribute_id"`
	Name                string             `bson:"name"`
	IsPreSelected       bool               `bson:"is_pre_selected"`
	DisplayOrder        int                `bson:"display_order"`
}

type CustomerAttributeValueRepository interface {
	Create(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Update(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Delete(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Fetch(c context.Context) ([]CustomerAttributeValue, error)
	FetchByID(c context.Context, customer_attribute_valueID string) (CustomerAttributeValue, error)
}

type CustomerAttributeValueUsecase interface {
	FetchByID(c context.Context, customer_attribute_valueID string) (CustomerAttributeValue, error)
	Create(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Update(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Delete(c context.Context, customer_attribute_value *CustomerAttributeValue) error
	Fetch(c context.Context) ([]CustomerAttributeValue, error)
}
