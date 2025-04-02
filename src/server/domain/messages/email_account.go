package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionEmailAccount = "email_accounts"
)

// EmailAccount represent an email account
type EmailAccount struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	Email                       string             `bson:"email"`
	DisplayName                 string             `bson:"display_name"`
	Host                        string             `bson:"host"`
	Port                        int                `bson:"port"`
	Username                    string             `bson:"username"`
	Password                    string             `bson:"password"`
	EnableSsl                   bool               `bson:"enable_ssl"`
	MaxNumberOfEmails           int                `bson:"max_number_of_emails"`
	EmailAuthenticationMethodID int                `bson:"email_authentication_method_id"`
	ClientID                    string             `bson:"client_id"`
	ClientSecret                string             `bson:"client_secret"`
	TenantID                    string             `bson:"tenant_id"`
	EmailAuthenticationMethod   int                `bson:"email_authentication_method"`
}

// EmailAccountRepository represents the repository interface for EmailAccount
type EmailAccountRepository interface {
	CreateMany(c context.Context, items []EmailAccount) error
	Create(c context.Context, email_account *EmailAccount) error
	Update(c context.Context, email_account *EmailAccount) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]EmailAccount, error)
	FetchByID(c context.Context, ID string) (EmailAccount, error)
}

// EmailAccountUsecase represents the use case interface for EmailAccount
type EmailAccountUsecase interface {
	CreateMany(c context.Context, items []EmailAccount) error
	FetchByID(c context.Context, ID string) (EmailAccount, error)
	Create(c context.Context, email_account *EmailAccount) error
	Update(c context.Context, email_account *EmailAccount) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]EmailAccount, error)
}
