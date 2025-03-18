package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type rewardpointshistoryUsecase struct {
	rewardpointshistoryRepository domain.RewardPointsHistoryRepository
	contextTimeout                time.Duration
}

func NewRewardPointsHistoryUsecase(rewardpointshistoryRepository domain.RewardPointsHistoryRepository, timeout time.Duration) domain.RewardPointsHistoryUsecase {
	return &rewardpointshistoryUsecase{
		rewardpointshistoryRepository: rewardpointshistoryRepository,
		contextTimeout:                timeout,
	}
}

func (tu *rewardpointshistoryUsecase) Create(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointshistoryRepository.Create(ctx, rewardpointshistory)
}

func (tu *rewardpointshistoryUsecase) Update(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointshistoryRepository.Update(ctx, rewardpointshistory)
}

func (tu *rewardpointshistoryUsecase) Delete(c context.Context, rewardpointshistory *domain.RewardPointsHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.rewardpointshistoryRepository.Delete(ctx, rewardpointshistory)
}

func (lu *rewardpointshistoryUsecase) FetchByID(c context.Context, rewardpointshistoryID string) (domain.RewardPointsHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.rewardpointshistoryRepository.FetchByID(ctx, rewardpointshistoryID)
}

func (lu *rewardpointshistoryUsecase) Fetch(c context.Context) ([]domain.RewardPointsHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.rewardpointshistoryRepository.Fetch(ctx)
}
