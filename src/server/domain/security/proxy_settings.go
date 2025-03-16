package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionProxySettings = "proxy_settings"
)

// ProxySettings represents proxy settings
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
