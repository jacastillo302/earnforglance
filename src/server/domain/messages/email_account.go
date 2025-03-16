package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionEmailAccount = "email_accounts"
)

// EmailAccount represents an email account
type EmailAccount struct {
	ID                          primitive.ObjectID        `bson:"_id,omitempty"`
	Email                       string                    `bson:"email"`
	DisplayName                 string                    `bson:"display_name"`
	Host                        string                    `bson:"host"`
	Port                        int                       `bson:"port"`
	Username                    string                    `bson:"username"`
	Password                    string                    `bson:"password"`
	EnableSsl                   bool                      `bson:"enable_ssl"`
	MaxNumberOfEmails           int                       `bson:"max_number_of_emails"`
	EmailAuthenticationMethodID int                       `bson:"email_authentication_method_id"`
	ClientID                    string                    `bson:"client_id"`
	ClientSecret                string                    `bson:"client_secret"`
	TenantID                    string                    `bson:"tenant_id"`
	EmailAuthenticationMethod   EmailAuthenticationMethod `bson:"email_authentication_method"`
}
