package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionDeliveryDate = "delivery_dates"
)

// DeliveryDate represents a delivery date.
type DeliveryDate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	DisplayOrder int                `bson:"display_order"`
}

// DeliveryDateRepository defines the repository interface for DeliveryDate
type DeliveryDateRepository interface {
	Create(c context.Context, delivery_date *DeliveryDate) error
	Update(c context.Context, delivery_date *DeliveryDate) error
	Delete(c context.Context, delivery_date *DeliveryDate) error
	Fetch(c context.Context) ([]DeliveryDate, error)
	FetchByID(c context.Context, delivery_dateID string) (DeliveryDate, error)
}

// DeliveryDateUsecase defines the use case interface for DeliveryDate
type DeliveryDateUsecase interface {
	FetchByID(c context.Context, delivery_dateID string) (DeliveryDate, error)
	Create(c context.Context, delivery_date *DeliveryDate) error
	Update(c context.Context, delivery_date *DeliveryDate) error
	Delete(c context.Context, delivery_date *DeliveryDate) error
	Fetch(c context.Context) ([]DeliveryDate, error)
}
