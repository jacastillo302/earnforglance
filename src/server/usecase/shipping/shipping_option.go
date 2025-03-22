package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shippingoptionUsecase struct {
	shippingoptionRepository domain.ShippingOptionRepository
	contextTimeout           time.Duration
}

func NewShippingOptionUsecase(shippingoptionRepository domain.ShippingOptionRepository, timeout time.Duration) domain.ShippingOptionUsecase {
	return &shippingoptionUsecase{
		shippingoptionRepository: shippingoptionRepository,
		contextTimeout:           timeout,
	}
}

func (tu *shippingoptionUsecase) Create(c context.Context, shippingoption *domain.ShippingOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingoptionRepository.Create(ctx, shippingoption)
}

func (tu *shippingoptionUsecase) Update(c context.Context, shippingoption *domain.ShippingOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingoptionRepository.Update(ctx, shippingoption)
}

func (tu *shippingoptionUsecase) Delete(c context.Context, shippingoption string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingoptionRepository.Delete(ctx, shippingoption)
}

func (lu *shippingoptionUsecase) FetchByID(c context.Context, shippingoptionID string) (domain.ShippingOption, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingoptionRepository.FetchByID(ctx, shippingoptionID)
}

func (lu *shippingoptionUsecase) Fetch(c context.Context) ([]domain.ShippingOption, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingoptionRepository.Fetch(ctx)
}
