package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPollVotingRecord = "poll_voting_records"
)

// PollVotingRecord represents a poll voting record
type PollVotingRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PollAnswerID primitive.ObjectID `bson:"poll_answer_id"`
	CustomerID   primitive.ObjectID `bson:"customer_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}

// PollVotingRecordRepository defines the repository interface for PollVotingRecord
type PollVotingRecordRepository interface {
	Create(c context.Context, poll_voting_record *PollVotingRecord) error
	Update(c context.Context, poll_voting_record *PollVotingRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollVotingRecord, error)
	FetchByID(c context.Context, ID string) (PollVotingRecord, error)
}

// PollVotingRecordUsecase defines the usecase interface for PollVotingRecord
type PollVotingRecordUsecase interface {
	FetchByID(c context.Context, ID string) (PollVotingRecord, error)
	Create(c context.Context, poll_voting_record *PollVotingRecord) error
	Update(c context.Context, poll_voting_record *PollVotingRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollVotingRecord, error)
}
