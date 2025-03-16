package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionStateProvince = "state_provinces"
)

// StateProvince represents a state/province
type StateProvince struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CountryID    int                `bson:"country_id"`
	Name         string             `bson:"name"`
	Abbreviation string             `bson:"abbreviation"`
	Published    bool               `bson:"published"`
	DisplayOrder int                `bson:"display_order"`
}
