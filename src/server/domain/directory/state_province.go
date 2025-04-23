package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionStateProvince = "state_provinces"
)

// StateProvince represents a state/province
type StateProvince struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	CountryID    bson.ObjectID `bson:"country_id"`
	Name         string        `bson:"name"`
	Abbreviation string        `bson:"abbreviation"`
	Published    bool          `bson:"published"`
	DisplayOrder int           `bson:"display_order"`
}

type StateProvinceRepository interface {
	CreateMany(c context.Context, items []StateProvince) error
	Create(c context.Context, state_province *StateProvince) error
	Update(c context.Context, state_province *StateProvince) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StateProvince, error)
	FetchByID(c context.Context, ID string) (StateProvince, error)
}

type StateProvinceUsecase interface {
	CreateMany(c context.Context, items []StateProvince) error
	FetchByID(c context.Context, ID string) (StateProvince, error)
	Create(c context.Context, state_province *StateProvince) error
	Update(c context.Context, state_province *StateProvince) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StateProvince, error)
}
