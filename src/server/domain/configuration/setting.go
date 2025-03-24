package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSetting = "settings"
)

// Setting represents a setting
type Setting struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Value   string             `bson:"value"`
	StoreID primitive.ObjectID `bson:"store_id"`
}

type SettingRepository interface {
	CreateMany(c context.Context, items []Setting) error
	Create(c context.Context, setting *Setting) error
	Update(c context.Context, setting *Setting) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Setting, error)
	FetchByID(c context.Context, ID string) (Setting, error)
}

type SettingUsecase interface {
	CreateMany(c context.Context, items []Setting) error
	FetchByID(c context.Context, ID string) (Setting, error)
	Create(c context.Context, setting *Setting) error
	Update(c context.Context, setting *Setting) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Setting, error)
}

// NewSetting creates a new Setting instance
func NewSetting(name string, value string, storeID primitive.ObjectID) *Setting {
	return &Setting{
		Name:    name,
		Value:   value,
		StoreID: storeID,
	}
}

// ToString returns the name of the setting
func (s *Setting) ToString() string {
	return s.Name
}
