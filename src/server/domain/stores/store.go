package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectioStore = "stores"
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
	DefaultLanguageID      int                `bson:"default_language_id"`
	DisplayOrder           int                `bson:"display_order"`
	CompanyName            string             `bson:"company_name"`
	CompanyAddress         string             `bson:"company_address"`
	CompanyPhoneNumber     string             `bson:"company_phone_number"`
	CompanyVat             string             `bson:"company_vat"`
	Deleted                bool               `bson:"deleted"`
}
