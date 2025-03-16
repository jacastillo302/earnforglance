package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

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
