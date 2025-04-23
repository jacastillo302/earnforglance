package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionAffiliate = "affiliates"
)

// Affiliate represents an affiliate
type Affiliate struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	AddressID       bson.ObjectID `bson:"address_id"`
	AdminComment    string        `bson:"admin_comment"`
	FriendlyUrlName string        `bson:"friendly_url_name"`
	Deleted         bool          `bson:"deleted"`
	Active          bool          `bson:"active"`
}

type AffiliateRepository interface {
	CreateMany(c context.Context, items []Affiliate) error
	FetchByID(c context.Context, ID string) (Affiliate, error)
	Create(c context.Context, affiliate *Affiliate) error
	Update(c context.Context, affiliate *Affiliate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Affiliate, error)
}

type AffiliateUsecase interface {
	CreateMany(c context.Context, items []Affiliate) error
	FetchByID(c context.Context, ID string) (Affiliate, error)
	Create(c context.Context, affiliate *Affiliate) error
	Update(c context.Context, affiliate *Affiliate) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Affiliate, error)
}
