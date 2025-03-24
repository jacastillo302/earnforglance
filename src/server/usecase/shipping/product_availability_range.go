package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type productavailabilityrangeUsecase struct {
	productavailabilityrangeRepository domain.ProductAvailabilityRangeRepository
	contextTimeout                     time.Duration
}

func NewProductAvailabilityRangeUsecase(productavailabilityrangeRepository domain.ProductAvailabilityRangeRepository, timeout time.Duration) domain.ProductAvailabilityRangeUsecase {
	return &productavailabilityrangeUsecase{
		productavailabilityrangeRepository: productavailabilityrangeRepository,
		contextTimeout:                     timeout,
	}
}

func (tu *productavailabilityrangeUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.ProductAvailabilityRange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productavailabilityrangeRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *productavailabilityrangeUsecase) Create(c context.Context, productavailabilityrange *domain.ProductAvailabilityRange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productavailabilityrangeRepository.Create(ctx, productavailabilityrange)
}

func (tu *productavailabilityrangeUsecase) Update(c context.Context, productavailabilityrange *domain.ProductAvailabilityRange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productavailabilityrangeRepository.Update(ctx, productavailabilityrange)
}

func (tu *productavailabilityrangeUsecase) Delete(c context.Context, productavailabilityrange string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productavailabilityrangeRepository.Delete(ctx, productavailabilityrange)
}

func (lu *productavailabilityrangeUsecase) FetchByID(c context.Context, productavailabilityrangeID string) (domain.ProductAvailabilityRange, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productavailabilityrangeRepository.FetchByID(ctx, productavailabilityrangeID)
}

func (lu *productavailabilityrangeUsecase) Fetch(c context.Context) ([]domain.ProductAvailabilityRange, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productavailabilityrangeRepository.Fetch(ctx)
}
