package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionMeasureSettings = "measure_settings"
)

// MeasureSettings represents measure settings
type MeasureSettings struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	BaseDimensionID int                `bson:"base_dimension_id"`
	BaseWeightID    int                `bson:"base_weight_id"`
}

type MeasureSettingsRepository interface {
	Create(c context.Context, measure_settings *MeasureSettings) error
	Update(c context.Context, measure_settings *MeasureSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureSettings, error)
	FetchByID(c context.Context, ID string) (MeasureSettings, error)
}

type MeasureSettingsUsecase interface {
	FetchByID(c context.Context, ID string) (MeasureSettings, error)
	Create(c context.Context, measure_settings *MeasureSettings) error
	Update(c context.Context, measure_settings *MeasureSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureSettings, error)
}
