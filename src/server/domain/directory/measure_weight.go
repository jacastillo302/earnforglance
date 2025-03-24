package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMeasureWeight = "measure_weights"
)

// MeasureWeight represents a measure weight
type MeasureWeight struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	SystemKeyword string             `bson:"system_keyword"`
	Ratio         float64            `bson:"ratio"`
	DisplayOrder  int                `bson:"display_order"`
}

type MeasureWeightRepository interface {
	CreateMany(c context.Context, items []MeasureWeight) error
	Create(c context.Context, measure_weight *MeasureWeight) error
	Update(c context.Context, measure_weight *MeasureWeight) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureWeight, error)
	FetchByID(c context.Context, ID string) (MeasureWeight, error)
}

type MeasureWeightUsecase interface {
	CreateMany(c context.Context, items []MeasureWeight) error
	FetchByID(c context.Context, ID string) (MeasureWeight, error)
	Create(c context.Context, measure_weight *MeasureWeight) error
	Update(c context.Context, measure_weight *MeasureWeight) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureWeight, error)
}
