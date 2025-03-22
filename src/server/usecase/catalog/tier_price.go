package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type tierpriceUsecase struct {
	tierpriceRepository domain.TierPriceRepository
	contextTimeout      time.Duration
}

func NewTierPriceUsecase(tierpriceRepository domain.TierPriceRepository, timeout time.Duration) domain.TierPriceUsecase {
	return &tierpriceUsecase{
		tierpriceRepository: tierpriceRepository,
		contextTimeout:      timeout,
	}
}

func (tu *tierpriceUsecase) Create(c context.Context, tierprice *domain.TierPrice) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tierpriceRepository.Create(ctx, tierprice)
}

func (tu *tierpriceUsecase) Update(c context.Context, tierprice *domain.TierPrice) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tierpriceRepository.Update(ctx, tierprice)
}

func (tu *tierpriceUsecase) Delete(c context.Context, tierprice string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.tierpriceRepository.Delete(ctx, tierprice)
}

func (lu *tierpriceUsecase) FetchByID(c context.Context, tierpriceID string) (domain.TierPrice, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.tierpriceRepository.FetchByID(ctx, tierpriceID)
}

func (lu *tierpriceUsecase) Fetch(c context.Context) ([]domain.TierPrice, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.tierpriceRepository.Fetch(ctx)
}
