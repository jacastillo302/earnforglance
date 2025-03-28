package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionApiClient = "api_settiings"
)

// ApiClient represents the app data entity
type ApiClient struct {
	ID                            primitive.ObjectID `bson:"_id,omitempty"`                    // MongoDB document ID                      string             `bson:"client_id"`                        // Represents the client ID
	Secret                        string             `bson:"secret_api"`                       // Represents the secret API key
	Enable                        bool               `bson:"enable_api"`                       // Indicates if the API is enabled
	DateExpired                   time.Time          `bson:"date_expired_api"`                 // Represents the API expiration date
	Name                          string             `bson:"name"`                             // Represents the name
	IdentityUrl                   string             `bson:"identity_url"`                     // Represents the identity URL
	CallbackUrl                   string             `bson:"callback_url"`                     // Represents the callback URL
	DefaultAccessTokenExpiration  int                `bson:"default_access_token_expiration"`  // Represents the default access token expiration time
	DefaultRefreshTokenExpiration int                `bson:"default_refresh_token_expiration"` // Represents the default refresh token expiration time
	CreatedDate                   time.Time          `bson:"created_date"`                     // Represents the creation date
}

type ApiClientRepository interface {
	CreateMany(c context.Context, items []ApiClient) error
	Create(c context.Context, affiliate *ApiClient) error
	Update(c context.Context, affiliate *ApiClient) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ApiClient, error)
	FetchByID(c context.Context, ID string) (ApiClient, error)
}

type ApiClientUsecase interface {
	CreateMany(c context.Context, items []ApiClient) error
	Create(c context.Context, affiliate *ApiClient) error
	Update(c context.Context, affiliate *ApiClient) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ApiClient, error)
	FetchByID(c context.Context, ID string) (ApiClient, error)
}
