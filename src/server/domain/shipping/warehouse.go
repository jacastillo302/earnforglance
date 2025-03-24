package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionWarehouse = "warehouses"
)

// Warehouse represents a warehouse
type Warehouse struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	AdminComment string             `bson:"admin_comment"`
	AddressID    primitive.ObjectID `bson:"address_id"`
}

type WarehouseRepository interface {
	CreateMany(c context.Context, items []Warehouse) error
	Create(c context.Context, warehouse *Warehouse) error
	Update(c context.Context, warehouse *Warehouse) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Warehouse, error)
	FetchByID(c context.Context, ID string) (Warehouse, error)
}

type WarehouseUsecase interface {
	FetchByID(c context.Context, ID string) (Warehouse, error)
	Create(c context.Context, warehouse *Warehouse) error
	Update(c context.Context, warehouse *Warehouse) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Warehouse, error)
}
