package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionExternalAuthenticationRecord = "external_authentication_record"
)

// ExternalAuthenticationRecord represents an external authentication record.
type ExternalAuthenticationRecord struct {
	ID                        bson.ObjectID `bson:"_id,omitempty"`               // MongoDB ObjectID
	CustomerID                bson.ObjectID `bson:"customer_id"`                 // Customer identifier
	Email                     string        `bson:"email"`                       // External email
	ExternalIdentifier        string        `bson:"external_identifier"`         // External identifier
	ExternalDisplayIdentifier string        `bson:"external_display_identifier"` // External display identifier
	OAuthToken                string        `bson:"oauth_token"`                 // OAuth token
	OAuthAccessToken          string        `bson:"oauth_access_token"`          // OAuth access token
	ProviderSystemName        string        `bson:"provider_system_name"`        // Provider system name
}

type ExternalAuthenticationRecordRepository interface {
	CreateMany(c context.Context, items []ExternalAuthenticationRecord) error
	Create(c context.Context, permission_record_customer_role_mapping *ExternalAuthenticationRecord) error
	Update(c context.Context, permission_record_customer_role_mapping *ExternalAuthenticationRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ExternalAuthenticationRecord, error)
	FetchByID(c context.Context, ID string) (ExternalAuthenticationRecord, error)
}

type ExternalAuthenticationRecordUsecase interface {
	CreateMany(c context.Context, items []ExternalAuthenticationRecord) error
	FetchByID(c context.Context, ID string) (ExternalAuthenticationRecord, error)
	Create(c context.Context, permission_record_customer_role_mapping *ExternalAuthenticationRecord) error
	Update(c context.Context, permission_record_customer_role_mapping *ExternalAuthenticationRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]ExternalAuthenticationRecord, error)
}
