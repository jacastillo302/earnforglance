package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPollAnswer = "poll_answers"
)

// PollAnswer represents a poll answer
type PollAnswer struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	PollID        bson.ObjectID `bson:"poll_id"`
	Name          string        `bson:"name"`
	NumberOfVotes int           `bson:"number_of_votes"`
	DisplayOrder  int           `bson:"display_order"`
}

// PollAnswerRepository defines the repository interface for PollAnswer
type PollAnswerRepository interface {
	CreateMany(c context.Context, items []PollAnswer) error
	Create(c context.Context, poll_answer *PollAnswer) error
	Update(c context.Context, poll_answer *PollAnswer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollAnswer, error)
	FetchByID(c context.Context, ID string) (PollAnswer, error)
}

// PollAnswerUsecase defines the usecase interface for PollAnswer
type PollAnswerUsecase interface {
	CreateMany(c context.Context, items []PollAnswer) error
	FetchByID(c context.Context, ID string) (PollAnswer, error)
	Create(c context.Context, poll_answer *PollAnswer) error
	Update(c context.Context, poll_answer *PollAnswer) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollAnswer, error)
}
