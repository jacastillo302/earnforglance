package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionPollVotingRecord = "poll_voting_records"
)

// PollVotingRecord represents a poll voting record
type PollVotingRecord struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	PollAnswerID bson.ObjectID `bson:"poll_answer_id"`
	CustomerID   bson.ObjectID `bson:"customer_id"`
	CreatedOnUtc time.Time     `bson:"created_on_utc"`
}

// PollVotingRecordRepository defines the repository interface for PollVotingRecord
type PollVotingRecordRepository interface {
	CreateMany(c context.Context, items []PollVotingRecord) error
	Create(c context.Context, poll_voting_record *PollVotingRecord) error
	Update(c context.Context, poll_voting_record *PollVotingRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollVotingRecord, error)
	FetchByID(c context.Context, ID string) (PollVotingRecord, error)
}

// PollVotingRecordUsecase defines the usecase interface for PollVotingRecord
type PollVotingRecordUsecase interface {
	CreateMany(c context.Context, items []PollVotingRecord) error
	FetchByID(c context.Context, ID string) (PollVotingRecord, error)
	Create(c context.Context, poll_voting_record *PollVotingRecord) error
	Update(c context.Context, poll_voting_record *PollVotingRecord) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]PollVotingRecord, error)
}
