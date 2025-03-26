package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionExternalAuthenticationRecord = "external_authentication_record"
)

// ExternalAuthenticationRecord represents an external authentication record.
type ExternalAuthenticationRecord struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`               // MongoDB ObjectID
	CustomerID                primitive.ObjectID `bson:"customer_id"`                 // Customer identifier
	Email                     string             `bson:"email"`                       // External email
	ExternalIdentifier        string             `bson:"external_identifier"`         // External identifier
	ExternalDisplayIdentifier string             `bson:"external_display_identifier"` // External display identifier
	OAuthToken                string             `bson:"oauth_token"`                 // OAuth token
	OAuthAccessToken          string             `bson:"oauth_access_token"`          // OAuth access token
	ProviderSystemName        string             `bson:"provider_system_name"`        // Provider system name
}
