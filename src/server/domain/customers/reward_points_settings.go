package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CollectionRewardPointsSettings = "reward_points_settings"
)

// RewardPointsSettings represents reward points settings
type RewardPointsSettings struct {
	ID                               bson.ObjectID `bson:"_id,omitempty"`
	Enabled                          bool          `bson:"enabled"`
	ExchangeRate                     float64       `bson:"exchange_rate"`
	MinimumRewardPointsToUse         int           `bson:"minimum_reward_points_to_use"`
	MaximumRewardPointsToUsePerOrder int           `bson:"maximum_reward_points_to_use_per_order"`
	MaximumRedeemedRate              float64       `bson:"maximum_redeemed_rate"`
	PointsForRegistration            int           `bson:"points_for_registration"`
	RegistrationPointsValidity       *int          `bson:"registration_points_validity"`
	PointsForPurchasesAmount         float64       `bson:"points_for_purchases_amount"`
	PointsForPurchasesPoints         int           `bson:"points_for_purchases_points"`
	PurchasesPointsValidity          *int          `bson:"purchases_points_validity"`
	MinOrderTotalToAwardPoints       float64       `bson:"min_order_total_to_award_points"`
	ActivationDelay                  int           `bson:"activation_delay"`
	ActivationDelayPeriodID          int           `bson:"activation_delay_period_id"`
	DisplayHowMuchWillBeEarned       bool          `bson:"display_how_much_will_be_earned"`
	PointsAccumulatedForAllStores    bool          `bson:"points_accumulated_for_all_stores"`
	PageSize                         int           `bson:"page_size"`
}

type RewardPointsSettingsRepository interface {
	CreateMany(c context.Context, items []RewardPointsSettings) error
	Create(c context.Context, reward_points_settings *RewardPointsSettings) error
	Update(c context.Context, reward_points_settings *RewardPointsSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RewardPointsSettings, error)
	FetchByID(c context.Context, ID string) (RewardPointsSettings, error)
}

type RewardPointsSettingsUsecase interface {
	CreateMany(c context.Context, items []RewardPointsSettings) error
	FetchByID(c context.Context, ID string) (RewardPointsSettings, error)
	Create(c context.Context, reward_points_settings *RewardPointsSettings) error
	Update(c context.Context, reward_points_settings *RewardPointsSettings) error
	Delete(c context.Context, ID string) error
	Fetch(c context.Context) ([]RewardPointsSettings, error)
}
