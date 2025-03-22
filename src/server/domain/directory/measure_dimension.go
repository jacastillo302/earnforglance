package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMeasureDimension = "measure_dimensions"
)

// MeasureDimension represents a measure dimension
type MeasureDimension struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	SystemKeyword string             `bson:"system_keyword"`
	Ratio         float64            `bson:"ratio"`
	DisplayOrder  int                `bson:"display_order"`
}

type MeasureDimensionRepository interface {
	Create(c context.Context, measure_dimension *MeasureDimension) error
	Update(c context.Context, measure_dimension *MeasureDimension) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureDimension, error)
	FetchByID(c context.Context, ID string) (MeasureDimension, error)
}

type MeasureDimensionUsecase interface {
	FetchByID(c context.Context, ID string) (MeasureDimension, error)
	Create(c context.Context, measure_dimension *MeasureDimension) error
	Update(c context.Context, measure_dimension *MeasureDimension) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureDimension, error)
}
