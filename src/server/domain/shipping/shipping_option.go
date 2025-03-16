package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionShippingOption = "shipping_options"
)

// ShippingOption represents a shipping option
type ShippingOption struct {
	ID                                      primitive.ObjectID `bson:"_id,omitempty"`
	ShippingRateComputationMethodSystemName string             `bson:"shipping_rate_computation_method_system_name"`
	Rate                                    float64            `bson:"rate"`
	Name                                    string             `bson:"name"`
	Description                             string             `bson:"description"`
	TransitDays                             *int               `bson:"transit_days,omitempty"`
	IsPickupInStore                         bool               `bson:"is_pickup_in_store"`
	DisplayOrder                            *int               `bson:"display_order,omitempty"`
}
