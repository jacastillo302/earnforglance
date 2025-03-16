package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionEmailAccountSettings = "email_account_settings"
)

// EmailAccountSettings represents email account settings
type EmailAccountSettings struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	DefaultEmailAccountID int                `bson:"default_email_account_id"`
}
