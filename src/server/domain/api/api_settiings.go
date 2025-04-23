package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionApiSettings = "api_settiings"
)

// ApiSettings represents the API settings configuration
type ApiSettings struct {
	ID                       bson.ObjectID `bson:"_id,omitempty"`               // MongoDB document ID
	EnableApi                bool          `bson:"enable_api"`                  // Indicates if the API is enabled
	AllowRequestsFromSwagger bool          `bson:"allow_requests_from_swagger"` // Indicates if requests from Swagger are allowed
	EnableLogging            bool          `bson:"enable_logging"`              // Indicates if logging is enabled
	TokenKey                 string        `bson:"token_key"`                   // The token key for API authentication
}

type ApiSettingsRepository interface {
	CreateMany(c context.Context, items []ApiSettings) error
	Create(c context.Context, item *ApiSettings) error
	Update(c context.Context, item *ApiSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ApiSettings, error)
	FetchByID(c context.Context, ID string) (ApiSettings, error)
}

type ApiSettingsUsecase interface {
	CreateMany(c context.Context, items []ApiSettings) error
	Create(c context.Context, item *ApiSettings) error
	Update(c context.Context, item *ApiSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ApiSettings, error)
	FetchByID(c context.Context, ID string) (ApiSettings, error)
}
