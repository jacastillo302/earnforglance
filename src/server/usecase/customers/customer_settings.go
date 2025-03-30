package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customersettingsUsecase struct {
	customersettingsRepository domain.CustomerSettingsRepository
	contextTimeout             time.Duration
}

func NewCustomerSettingsUsecase(customersettingsRepository domain.CustomerSettingsRepository, timeout time.Duration) domain.CustomerSettingsUsecase {
	return &customersettingsUsecase{
		customersettingsRepository: customersettingsRepository,
		contextTimeout:             timeout,
	}
}

func (tu *customersettingsUsecase) CreateMany(c context.Context, items []domain.CustomerSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customersettingsRepository.CreateMany(ctx, items)
}

func (tu *customersettingsUsecase) Create(c context.Context, customersettings *domain.CustomerSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customersettingsRepository.Create(ctx, customersettings)
}

func (tu *customersettingsUsecase) Update(c context.Context, customersettings *domain.CustomerSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customersettingsRepository.Update(ctx, customersettings)
}

func (tu *customersettingsUsecase) Delete(c context.Context, customersettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customersettingsRepository.Delete(ctx, customersettings)
}

func (lu *customersettingsUsecase) FetchByID(c context.Context, customersettingsID string) (domain.CustomerSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customersettingsRepository.FetchByID(ctx, customersettingsID)
}

func (lu *customersettingsUsecase) Fetch(c context.Context) ([]domain.CustomerSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customersettingsRepository.Fetch(ctx)
}
