package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type orderSettingsUsecase struct {
	orderSettingsRepository domain.OrderSettingsRepository
	contextTimeout          time.Duration
}

func NewOrderSettingsUsecase(orderSettingsRepository domain.OrderSettingsRepository, timeout time.Duration) domain.OrderSettingsUsecase {
	return &orderSettingsUsecase{
		orderSettingsRepository: orderSettingsRepository,
		contextTimeout:          timeout,
	}
}

func (tu *orderSettingsUsecase) Create(c context.Context, orderSettings *domain.OrderSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderSettingsRepository.Create(ctx, orderSettings)
}

func (tu *orderSettingsUsecase) Update(c context.Context, orderSettings *domain.OrderSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderSettingsRepository.Update(ctx, orderSettings)
}

func (tu *orderSettingsUsecase) Delete(c context.Context, orderSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderSettingsRepository.Delete(ctx, orderSettings)
}

func (lu *orderSettingsUsecase) FetchByID(c context.Context, orderSettingsID string) (domain.OrderSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderSettingsRepository.FetchByID(ctx, orderSettingsID)
}

func (lu *orderSettingsUsecase) Fetch(c context.Context) ([]domain.OrderSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderSettingsRepository.Fetch(ctx)
}
