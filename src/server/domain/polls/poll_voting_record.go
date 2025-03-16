package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPollVotingRecord = "poll_voting_records"
)

// PollVotingRecord represents a poll voting record
type PollVotingRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PollAnswerID int                `bson:"poll_answer_id"`
	CustomerID   int                `bson:"customer_id"`
	CreatedOnUtc time.Time          `bson:"created_on_utc"`
}
