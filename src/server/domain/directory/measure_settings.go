package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionMeasureSettings = "measure_settings"
)

// MeasureSettings represents measure settings
type MeasureSettings struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	BaseDimensionID int                `bson:"base_dimension_id"`
	BaseWeightID    int                `bson:"base_weight_id"`
}
