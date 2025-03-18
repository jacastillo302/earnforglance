package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/discounts"
)

type discountManufacturerMappingUsecase struct {
	discountManufacturerMappingRepository domain.DiscountManufacturerMappingRepository
	contextTimeout                        time.Duration
}

func NewDiscountManufacturerMappingUsecase(discountManufacturerMappingRepository domain.DiscountManufacturerMappingRepository, timeout time.Duration) domain.DiscountManufacturerMappingUsecase {
	return &discountManufacturerMappingUsecase{
		discountManufacturerMappingRepository: discountManufacturerMappingRepository,
		contextTimeout:                        timeout,
	}
}

func (tu *discountManufacturerMappingUsecase) Create(c context.Context, discountManufacturerMapping *domain.DiscountManufacturerMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountManufacturerMappingRepository.Create(ctx, discountManufacturerMapping)
}

func (tu *discountManufacturerMappingUsecase) Update(c context.Context, discountManufacturerMapping *domain.DiscountManufacturerMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountManufacturerMappingRepository.Update(ctx, discountManufacturerMapping)
}

func (tu *discountManufacturerMappingUsecase) Delete(c context.Context, discountManufacturerMapping *domain.DiscountManufacturerMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.discountManufacturerMappingRepository.Delete(ctx, discountManufacturerMapping)
}

func (lu *discountManufacturerMappingUsecase) FetchByID(c context.Context, discountManufacturerMappingID string) (domain.DiscountManufacturerMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountManufacturerMappingRepository.FetchByID(ctx, discountManufacturerMappingID)
}

func (lu *discountManufacturerMappingUsecase) Fetch(c context.Context) ([]domain.DiscountManufacturerMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.discountManufacturerMappingRepository.Fetch(ctx)
}
