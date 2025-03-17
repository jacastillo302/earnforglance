package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAffiliate = "affiliates"
)

// Affiliate represents an affiliate
type Affiliate struct {
	AddressId       primitive.ObjectID `bson:"_id,omitempty"`
	AdminComment    string             `bson:"admin_comment"`
	FriendlyUrlName string             `bson:"friendly_url_name"`
	Deleted         bool               `bson:"deleted"`
	Active          bool               `bson:"active"`
}

type AffiliateRepository interface {
	Create(c context.Context, affiliate *Affiliate) error
	Update(c context.Context, affiliate *Affiliate) error
	Fetch(c context.Context) ([]Affiliate, error)
	GetActive(c context.Context, active bool) (Affiliate, error)
	FetchByID(c context.Context, affiliateID string) (Affiliate, error)
}

type AffiliateUsecase interface {
	FetchByID(c context.Context, affiliateID string) (Affiliate, error)
	Create(c context.Context, affiliate *Affiliate) error
	Update(c context.Context, affiliate *Affiliate) error
	Fetch(c context.Context) ([]Affiliate, error)
}
