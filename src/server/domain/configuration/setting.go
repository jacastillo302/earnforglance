package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSetting = "settings"
)

// Setting represents a setting
type Setting struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Value   string             `bson:"value"`
	StoreID int                `bson:"store_id"`
}

// NewSetting creates a new Setting instance
func NewSetting(name string, value string, storeID int) *Setting {
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
