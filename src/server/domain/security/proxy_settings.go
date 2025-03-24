package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProxySettings = "proxy_settings"
)

// ProxySettings represents proxy settings.
type ProxySettings struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Enabled         bool               `bson:"enabled"`
	Address         string             `bson:"address"`
	Port            string             `bson:"port"`
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	BypassOnLocal   bool               `bson:"bypass_on_local"`
	PreAuthenticate bool               `bson:"pre_authenticate"`
}

type ProxySettingsRepository interface {
	CreateMany(c context.Context, items []ProxySettings) error
	Create(c context.Context, proxy_settings *ProxySettings) error
	Update(c context.Context, proxy_settings *ProxySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProxySettings, error)
	FetchByID(c context.Context, ID string) (ProxySettings, error)
}

type ProxySettingsUsecase interface {
	CreateMany(c context.Context, items []ProxySettings) error
	FetchByID(c context.Context, ID string) (ProxySettings, error)
	Create(c context.Context, proxy_settings *ProxySettings) error
	Update(c context.Context, proxy_settings *ProxySettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ProxySettings, error)
}
