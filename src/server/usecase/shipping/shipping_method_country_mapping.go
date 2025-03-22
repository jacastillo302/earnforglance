package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shippingmethodcountrymappingUsecase struct {
	shippingmethodcountrymappingRepository domain.ShippingMethodCountryMappingRepository
	contextTimeout                         time.Duration
}

func NewShippingMethodCountryMappingUsecase(shippingmethodcountrymappingRepository domain.ShippingMethodCountryMappingRepository, timeout time.Duration) domain.ShippingMethodCountryMappingUsecase {
	return &shippingmethodcountrymappingUsecase{
		shippingmethodcountrymappingRepository: shippingmethodcountrymappingRepository,
		contextTimeout:                         timeout,
	}
}

func (tu *shippingmethodcountrymappingUsecase) Create(c context.Context, shippingmethodcountrymapping *domain.ShippingMethodCountryMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodcountrymappingRepository.Create(ctx, shippingmethodcountrymapping)
}

func (tu *shippingmethodcountrymappingUsecase) Update(c context.Context, shippingmethodcountrymapping *domain.ShippingMethodCountryMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodcountrymappingRepository.Update(ctx, shippingmethodcountrymapping)
}

func (tu *shippingmethodcountrymappingUsecase) Delete(c context.Context, shippingmethodcountrymapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodcountrymappingRepository.Delete(ctx, shippingmethodcountrymapping)
}

func (lu *shippingmethodcountrymappingUsecase) FetchByID(c context.Context, shippingmethodcountrymappingID string) (domain.ShippingMethodCountryMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingmethodcountrymappingRepository.FetchByID(ctx, shippingmethodcountrymappingID)
}

func (lu *shippingmethodcountrymappingUsecase) Fetch(c context.Context) ([]domain.ShippingMethodCountryMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingmethodcountrymappingRepository.Fetch(ctx)
}
