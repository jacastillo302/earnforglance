package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionDeliveryDate = "delivery_dates"
)

// DeliveryDate represents a delivery date.
type DeliveryDate struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	DisplayOrder int           `bson:"display_order"`
}

// DeliveryDateRepository defines the repository interface for DeliveryDate
type DeliveryDateRepository interface {
	CreateMany(c context.Context, items []DeliveryDate) error
	Create(c context.Context, delivery_date *DeliveryDate) error
	Update(c context.Context, delivery_date *DeliveryDate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DeliveryDate, error)
	FetchByID(c context.Context, ID string) (DeliveryDate, error)
}

// DeliveryDateUsecase defines the use case interface for DeliveryDate
type DeliveryDateUsecase interface {
	CreateMany(c context.Context, items []DeliveryDate) error
	FetchByID(c context.Context, ID string) (DeliveryDate, error)
	Create(c context.Context, delivery_date *DeliveryDate) error
	Update(c context.Context, delivery_date *DeliveryDate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]DeliveryDate, error)
}
