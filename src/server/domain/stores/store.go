package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStore = "stores"
)

// Store represents a store
type Store struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name"`
	DefaultMetaKeywords    string             `bson:"default_meta_keywords"`
	DefaultMetaDescription string             `bson:"default_meta_description"`
	DefaultTitle           string             `bson:"default_title"`
	HomepageTitle          string             `bson:"homepage_title"`
	HomepageDescription    string             `bson:"homepage_description"`
	Url                    string             `bson:"url"`
	SslEnabled             bool               `bson:"ssl_enabled"`
	Hosts                  string             `bson:"hosts"`
	DefaultLanguageID      primitive.ObjectID `bson:"default_language_id"`
	DisplayOrder           int                `bson:"display_order"`
	CompanyName            string             `bson:"company_name"`
	CompanyAddress         string             `bson:"company_address"`
	CompanyPhoneNumber     string             `bson:"company_phone_number"`
	CompanyVat             string             `bson:"company_vat"`
	Deleted                bool               `bson:"deleted"`
}

// StoreRepository defines the repository interface for Store
type StoreRepository interface {
	Create(c context.Context, store *Store) error
	Update(c context.Context, store *Store) error
	Delete(c context.Context, store *Store) error
	Fetch(c context.Context) ([]Store, error)
	FetchByID(c context.Context, storeID string) (Store, error)
}

// StoreUsecase defines the use case interface for Store
type StoreUsecase interface {
	FetchByID(c context.Context, storeID string) (Store, error)
	Create(c context.Context, store *Store) error
	Update(c context.Context, store *Store) error
	Delete(c context.Context, store *Store) error
	Fetch(c context.Context) ([]Store, error)
}
