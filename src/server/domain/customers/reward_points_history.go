package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionRewardPointsHistory = "reward_points_history"
)

// RewardPointsHistory represents a reward point history entry
type RewardPointsHistory struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	CustomerID    bson.ObjectID `bson:"customer_id"`
	StoreID       bson.ObjectID `bson:"store_id"`
	Points        int           `bson:"points"`
	PointsBalance *int          `bson:"points_balance"`
	UsedAmount    float64       `bson:"used_amount"`
	Message       string        `bson:"message"`
	CreatedOnUtc  time.Time     `bson:"created_on_utc"`
	EndDateUtc    *time.Time    `bson:"end_date_utc"`
	ValidPoints   *int          `bson:"valid_points"`
	UsedWithOrder *uuid.UUID    `bson:"used_with_order"`
}

type RewardPointsHistoryRepository interface {
	CreateMany(c context.Context, items []RewardPointsHistory) error
	Create(c context.Context, reward_point_history *RewardPointsHistory) error
	Update(c context.Context, reward_point_history *RewardPointsHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RewardPointsHistory, error)
	FetchByID(c context.Context, ID string) (RewardPointsHistory, error)
}

type RewardPointsHistoryUsecase interface {
	CreateMany(c context.Context, items []RewardPointsHistory) error
	FetchByID(c context.Context, ID string) (RewardPointsHistory, error)
	Create(c context.Context, reward_point_history *RewardPointsHistory) error
	Update(c context.Context, reward_point_history *RewardPointsHistory) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RewardPointsHistory, error)
}
