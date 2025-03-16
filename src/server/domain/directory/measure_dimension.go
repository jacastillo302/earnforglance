package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

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
