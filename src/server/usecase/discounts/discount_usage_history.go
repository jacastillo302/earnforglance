package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountusagehistoryUsecase struct {
	discountusagehistoryRepository domain.DiscountUsageHistoryRepository
	contextTimeout                 time.Duration
}

func NewDiscountUsageHistoryUsecase(discountusagehistoryRepository domain.DiscountUsageHistoryRepository, timeout time.Duration) domain.DiscountUsageHistoryUsecase {
	return &discountusagehistoryUsecase{
		discountusagehistoryRepository: discountusagehistoryRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *discountusagehistoryUsecase) Create(c context.Context, discountusagehistory *domain.DiscountUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountusagehistoryRepository.Create(ctx, discountusagehistory)
}

func (tu *discountusagehistoryUsecase) Update(c context.Context, discountusagehistory *domain.DiscountUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountusagehistoryRepository.Update(ctx, discountusagehistory)
}

func (tu *discountusagehistoryUsecase) Delete(c context.Context, discountusagehistory *domain.DiscountUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountusagehistoryRepository.Delete(ctx, discountusagehistory)
}

func (lu *discountusagehistoryUsecase) FetchByID(c context.Context, discountusagehistoryID string) (domain.DiscountUsageHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountusagehistoryRepository.FetchByID(ctx, discountusagehistoryID)
}

func (lu *discountusagehistoryUsecase) Fetch(c context.Context) ([]domain.DiscountUsageHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountusagehistoryRepository.Fetch(ctx)
}
