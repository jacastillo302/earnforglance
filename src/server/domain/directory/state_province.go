package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type StateProvinceRepository interface {
	Create(c context.Context, state_province *StateProvince) error
	Update(c context.Context, state_province *StateProvince) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StateProvince, error)
	FetchByID(c context.Context, ID string) (StateProvince, error)
}

type StateProvinceUsecase interface {
	FetchByID(c context.Context, ID string) (StateProvince, error)
	Create(c context.Context, state_province *StateProvince) error
	Update(c context.Context, state_province *StateProvince) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]StateProvince, error)
}
