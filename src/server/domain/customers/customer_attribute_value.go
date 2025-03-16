package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionCustomerAttributeValue = "customer_attribute_values"
)

// CustomerAttributeValue represents a customer attribute value
type CustomerAttributeValue struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	CustomerAttributeID int                `bson:"customer_attribute_id"`
	Name                string             `bson:"name"`
	IsPreSelected       bool               `bson:"is_pre_selected"`
	DisplayOrder        int                `bson:"display_order"`
}
