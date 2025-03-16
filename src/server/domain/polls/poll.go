package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPoll = "polls"
)

// Poll represents a poll
type Poll struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID        int                `bson:"language_id"`
	Name              string             `bson:"name"`
	SystemKeyword     string             `bson:"system_keyword"`
	Published         bool               `bson:"published"`
	ShowOnHomepage    bool               `bson:"show_on_homepage"`
	AllowGuestsToVote bool               `bson:"allow_guests_to_vote"`
	DisplayOrder      int                `bson:"display_order"`
	LimitedToStores   bool               `bson:"limited_to_stores"`
	StartDateUtc      *time.Time         `bson:"start_date_utc,omitempty"`
	EndDateUtc        *time.Time         `bson:"end_date_utc,omitempty"`
}
