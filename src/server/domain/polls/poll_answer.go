package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPollAnswer = "poll_answers"
)

// PollAnswer represents a poll answer
type PollAnswer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	PollID        primitive.ObjectID `bson:"poll_id"`
	Name          string             `bson:"name"`
	NumberOfVotes int                `bson:"number_of_votes"`
	DisplayOrder  int                `bson:"display_order"`
}

// PollAnswerRepository defines the repository interface for PollAnswer
type PollAnswerRepository interface {
	Create(c context.Context, poll_answer *PollAnswer) error
	Update(c context.Context, poll_answer *PollAnswer) error
	Delete(c context.Context, poll_answer *PollAnswer) error
	Fetch(c context.Context) ([]PollAnswer, error)
	FetchByID(c context.Context, poll_answerID string) (PollAnswer, error)
}

// PollAnswerUsecase defines the usecase interface for PollAnswer
type PollAnswerUsecase interface {
	FetchByID(c context.Context, poll_answerID string) (PollAnswer, error)
	Create(c context.Context, poll_answer *PollAnswer) error
	Update(c context.Context, poll_answer *PollAnswer) error
	Delete(c context.Context, poll_answer *PollAnswer) error
	Fetch(c context.Context) ([]PollAnswer, error)
}
