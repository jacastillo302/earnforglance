package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionSetting = "settings"
)

// Setting represents a setting
type Setting struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Name    string        `bson:"name"`
	Value   string        `bson:"value"`
	StoreID bson.ObjectID `bson:"store_id"`
}

type SettingRepository interface {
	CreateMany(c context.Context, items []Setting) error
	Create(c context.Context, setting *Setting) error
	Update(c context.Context, setting *Setting) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Setting, error)
	FetchByID(c context.Context, ID string) (Setting, error)
	FetchByName(c context.Context, name string) (Setting, error)
	FetchByNames(c context.Context, names []string) ([]Setting, error)
}

type SettingUsecase interface {
	CreateMany(c context.Context, items []Setting) error
	FetchByID(c context.Context, ID string) (Setting, error)
	Create(c context.Context, setting *Setting) error
	Update(c context.Context, setting *Setting) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Setting, error)
	FetchByName(c context.Context, name string) (Setting, error)
	FetchByNames(c context.Context, names []string) ([]Setting, error)
}

// NewSetting creates a new Setting instance
func NewSetting(name string, value string, storeID bson.ObjectID) *Setting {
	return &Setting{
		Name:    name,
		Value:   value,
		StoreID: storeID,
	}
}
