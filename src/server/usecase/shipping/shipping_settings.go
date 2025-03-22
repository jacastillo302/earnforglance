package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shippingSettingsUsecase struct {
	shippingSettingsRepository domain.ShippingSettingsRepository
	contextTimeout             time.Duration
}

func NewShippingSettingsUsecase(shippingSettingsRepository domain.ShippingSettingsRepository, timeout time.Duration) domain.ShippingSettingsUsecase {
	return &shippingSettingsUsecase{
		shippingSettingsRepository: shippingSettingsRepository,
		contextTimeout:             timeout,
	}
}

func (tu *shippingSettingsUsecase) Create(c context.Context, shippingSettings *domain.ShippingSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingSettingsRepository.Create(ctx, shippingSettings)
}

func (tu *shippingSettingsUsecase) Update(c context.Context, shippingSettings *domain.ShippingSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingSettingsRepository.Update(ctx, shippingSettings)
}

func (tu *shippingSettingsUsecase) Delete(c context.Context, shippingSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shippingSettingsRepository.Delete(ctx, shippingSettings)
}

func (lu *shippingSettingsUsecase) FetchByID(c context.Context, shippingSettingsID string) (domain.ShippingSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingSettingsRepository.FetchByID(ctx, shippingSettingsID)
}

func (lu *shippingSettingsUsecase) Fetch(c context.Context) ([]domain.ShippingSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shippingSettingsRepository.Fetch(ctx)
}
