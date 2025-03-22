package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountUsecase struct {
	discountRepository domain.DiscountRepository
	contextTimeout     time.Duration
}

func NewDiscountUsecase(discountRepository domain.DiscountRepository, timeout time.Duration) domain.DiscountUsecase {
	return &discountUsecase{
		discountRepository: discountRepository,
		contextTimeout:     timeout,
	}
}

func (tu *discountUsecase) Create(c context.Context, discount *domain.Discount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Create(ctx, discount)
}

func (tu *discountUsecase) Update(c context.Context, discount *domain.Discount) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Update(ctx, discount)
}

func (tu *discountUsecase) Delete(c context.Context, discount string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Delete(ctx, discount)
}

func (lu *discountUsecase) FetchByID(c context.Context, discountID string) (domain.Discount, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountRepository.FetchByID(ctx, discountID)
}

func (lu *discountUsecase) Fetch(c context.Context) ([]domain.Discount, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountRepository.Fetch(ctx)
}
