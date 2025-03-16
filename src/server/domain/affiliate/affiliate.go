package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionAffiliate = "affiliates"
)

// Affiliate represents an affiliate
type Affiliate struct {
	AddressId       primitive.ObjectID `bson:"address_id"`
	AdminComment    string             `bson:"admin_comment"`
	FriendlyUrlName string             `bson:"friendly_url_name"`
	Deleted         bool               `bson:"deleted"`
	Active          bool               `bson:"active"`
}
