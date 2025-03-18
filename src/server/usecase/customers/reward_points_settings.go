package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type rewardpointssettingsUsecase struct {
	rewardpointssettingsRepository domain.RewardPointsSettingsRepository
	contextTimeout                 time.Duration
}

func NewRewardPointsSettingsUsecase(rewardpointssettingsRepository domain.RewardPointsSettingsRepository, timeout time.Duration) domain.RewardPointsSettingsUsecase {
	return &rewardpointssettingsUsecase{
		rewardpointssettingsRepository: rewardpointssettingsRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *rewardpointssettingsUsecase) Create(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointssettingsRepository.Create(ctx, rewardpointssettings)
}

func (tu *rewardpointssettingsUsecase) Update(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointssettingsRepository.Update(ctx, rewardpointssettings)
}

func (tu *rewardpointssettingsUsecase) Delete(c context.Context, rewardpointssettings *domain.RewardPointsSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointssettingsRepository.Delete(ctx, rewardpointssettings)
}

func (lu *rewardpointssettingsUsecase) FetchByID(c context.Context, rewardpointssettingsID string) (domain.RewardPointsSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.rewardpointssettingsRepository.FetchByID(ctx, rewardpointssettingsID)
}

func (lu *rewardpointssettingsUsecase) Fetch(c context.Context) ([]domain.RewardPointsSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.rewardpointssettingsRepository.Fetch(ctx)
}
