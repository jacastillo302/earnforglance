package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shippingmethodUsecase struct {
	shippingmethodRepository domain.ShippingMethodRepository
	contextTimeout           time.Duration
}

func NewShippingMethodUsecase(shippingmethodRepository domain.ShippingMethodRepository, timeout time.Duration) domain.ShippingMethodUsecase {
	return &shippingmethodUsecase{
		shippingmethodRepository: shippingmethodRepository,
		contextTimeout:           timeout,
	}
}

func (tu *shippingmethodUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.ShippingMethod) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *shippingmethodUsecase) Create(c context.Context, shippingmethod *domain.ShippingMethod) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodRepository.Create(ctx, shippingmethod)
}

func (tu *shippingmethodUsecase) Update(c context.Context, shippingmethod *domain.ShippingMethod) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodRepository.Update(ctx, shippingmethod)
}

func (tu *shippingmethodUsecase) Delete(c context.Context, shippingmethod string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingmethodRepository.Delete(ctx, shippingmethod)
}

func (lu *shippingmethodUsecase) FetchByID(c context.Context, shippingmethodID string) (domain.ShippingMethod, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingmethodRepository.FetchByID(ctx, shippingmethodID)
}

func (lu *shippingmethodUsecase) Fetch(c context.Context) ([]domain.ShippingMethod, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingmethodRepository.Fetch(ctx)
}
