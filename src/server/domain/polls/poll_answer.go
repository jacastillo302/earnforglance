package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionPollAnswer = "poll_answers"
)

// PollAnswer represents a poll answer
type PollAnswer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	PollID        int                `bson:"poll_id"`
	Name          string             `bson:"name"`
	NumberOfVotes int                `bson:"number_of_votes"`
	DisplayOrder  int                `bson:"display_order"`
}
