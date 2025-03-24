package domain

import (
	"context" // Added context library
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPoll = "polls"
)

// Poll represents a poll
type Poll struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	LanguageID        primitive.ObjectID `bson:"language_id"`
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

// PollRepository defines the repository interface for Poll
type PollRepository interface {
	CreateMany(c context.Context, items []Poll) error
	Create(c context.Context, poll *Poll) error
	Update(c context.Context, poll *Poll) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Poll, error)
	FetchByID(c context.Context, ID string) (Poll, error)
}

// PollUsecase defines the use case interface for Poll
type PollUsecase interface {
	CreateMany(c context.Context, items []Poll) error
	FetchByID(c context.Context, ID string) (Poll, error)
	Create(c context.Context, poll *Poll) error
	Update(c context.Context, poll *Poll) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]Poll, error)
}
