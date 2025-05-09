package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionMeasureSettings = "measure_settings"
)

// MeasureSettings represents measure settings
type MeasureSettings struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	BaseDimensionID int           `bson:"base_dimension_id"`
	BaseWeightID    int           `bson:"base_weight_id"`
}

type MeasureSettingsRepository interface {
	CreateMany(c context.Context, items []MeasureSettings) error
	Create(c context.Context, measure_settings *MeasureSettings) error
	Update(c context.Context, measure_settings *MeasureSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureSettings, error)
	FetchByID(c context.Context, ID string) (MeasureSettings, error)
}

type MeasureSettingsUsecase interface {
	CreateMany(c context.Context, items []MeasureSettings) error
	FetchByID(c context.Context, ID string) (MeasureSettings, error)
	Create(c context.Context, measure_settings *MeasureSettings) error
	Update(c context.Context, measure_settings *MeasureSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]MeasureSettings, error)
}
