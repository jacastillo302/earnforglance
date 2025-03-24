package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRewardPointsHistory = "reward_points_history"
)

// RewardPointsHistory represents a reward point history entry
type RewardPointsHistory struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    primitive.ObjectID `bson:"customer_id"`
	StoreID       primitive.ObjectID `bson:"store_id"`
	Points        int                `bson:"points"`
	PointsBalance *int               `bson:"points_balance,omitempty"`
	UsedAmount    float64            `bson:"used_amount"`
	Message       string             `bson:"message"`
	CreatedOnUtc  time.Time          `bson:"created_on_utc"`
	EndDateUtc    *time.Time         `bson:"end_date_utc,omitempty"`
	ValidPoints   *int               `bson:"valid_points,omitempty"`
	UsedWithOrder *uuid.UUID         `bson:"used_with_order,omitempty"`
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
