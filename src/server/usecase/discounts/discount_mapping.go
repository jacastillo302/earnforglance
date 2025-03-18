package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountMappingUseCase struct {
	discountRepository domain.DiscountMappingRepository
	contextTimeout     time.Duration
}

func NewdiscountMappingUseCase(discountRepository domain.DiscountMappingRepository, timeout time.Duration) domain.DiscountMappingUsecase {
	return &discountMappingUseCase{
		discountRepository: discountRepository,
		contextTimeout:     timeout,
	}
}

func (tu *discountMappingUseCase) Create(c context.Context, discount *domain.DiscountMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Create(ctx, discount)
}

func (tu *discountMappingUseCase) Update(c context.Context, discount *domain.DiscountMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Update(ctx, discount)
}

func (tu *discountMappingUseCase) Delete(c context.Context, discount *domain.DiscountMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountRepository.Delete(ctx, discount)
}

func (lu *discountMappingUseCase) FetchByID(c context.Context, discountID string) (domain.DiscountMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountRepository.FetchByID(ctx, discountID)
}

func (lu *discountMappingUseCase) Fetch(c context.Context) ([]domain.DiscountMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountRepository.Fetch(ctx)
}
