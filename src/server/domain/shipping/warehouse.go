package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionWarehouse = "warehouses"
)

// Warehouse represents a warehouse
type Warehouse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	AdminComment string             `bson:"admin_comment"`
	AddressID    int                `bson:"address_id"`
}
