package domain

import (
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
	CustomerID    int                `bson:"customer_id"`
	StoreID       int                `bson:"store_id"`
	Points        int                `bson:"points"`
	PointsBalance *int               `bson:"points_balance,omitempty"`
	UsedAmount    float64            `bson:"used_amount"`
	Message       string             `bson:"message"`
	CreatedOnUtc  time.Time          `bson:"created_on_utc"`
	EndDateUtc    *time.Time         `bson:"end_date_utc,omitempty"`
	ValidPoints   *int               `bson:"valid_points,omitempty"`
	UsedWithOrder *uuid.UUID         `bson:"used_with_order,omitempty"`
}
