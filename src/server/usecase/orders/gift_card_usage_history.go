package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type giftcardusagehistoryUsecase struct {
	giftcardusagehistoryRepository domain.GiftCardUsageHistoryRepository
	contextTimeout                 time.Duration
}

func NewGiftCardUsageHistoryUsecase(giftcardusagehistoryRepository domain.GiftCardUsageHistoryRepository, timeout time.Duration) domain.GiftCardUsageHistoryUsecase {
	return &giftcardusagehistoryUsecase{
		giftcardusagehistoryRepository: giftcardusagehistoryRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *giftcardusagehistoryUsecase) Create(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardusagehistoryRepository.Create(ctx, giftcardusagehistory)
}

func (tu *giftcardusagehistoryUsecase) Update(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardusagehistoryRepository.Update(ctx, giftcardusagehistory)
}

func (tu *giftcardusagehistoryUsecase) Delete(c context.Context, giftcardusagehistory *domain.GiftCardUsageHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardusagehistoryRepository.Delete(ctx, giftcardusagehistory)
}

func (lu *giftcardusagehistoryUsecase) FetchByID(c context.Context, giftcardusagehistoryID string) (domain.GiftCardUsageHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.giftcardusagehistoryRepository.FetchByID(ctx, giftcardusagehistoryID)
}

func (lu *giftcardusagehistoryUsecase) Fetch(c context.Context) ([]domain.GiftCardUsageHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.giftcardusagehistoryRepository.Fetch(ctx)
}
