package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionWarehouse = "warehouses"
)

// Warehouse represents a warehouse
type Warehouse struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	AdminComment string        `bson:"admin_comment"`
	AddressID    bson.ObjectID `bson:"address_id"`
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
	CreateMany(c context.Context, items []Warehouse) error
	FetchByID(c context.Context, ID string) (Warehouse, error)
	Create(c context.Context, warehouse *Warehouse) error
	Update(c context.Context, warehouse *Warehouse) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Warehouse, error)
}
