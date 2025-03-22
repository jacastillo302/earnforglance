package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/tax"
)

type TaxSettingsUsecase struct {
	TaxSettingsRepository domain.TaxSettingsRepository
	contextTimeout        time.Duration
}

func NewTaxSettingsUsecase(TaxSettingsRepository domain.TaxSettingsRepository, timeout time.Duration) domain.TaxSettingsUsecase {
	return &TaxSettingsUsecase{
		TaxSettingsRepository: TaxSettingsRepository,
		contextTimeout:        timeout,
	}
}

func (tu *TaxSettingsUsecase) Create(c context.Context, TaxSettings *domain.TaxSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.TaxSettingsRepository.Create(ctx, TaxSettings)
}

func (tu *TaxSettingsUsecase) Update(c context.Context, TaxSettings *domain.TaxSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.TaxSettingsRepository.Update(ctx, TaxSettings)
}

func (tu *TaxSettingsUsecase) Delete(c context.Context, TaxSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.TaxSettingsRepository.Delete(ctx, TaxSettings)
}

func (lu *TaxSettingsUsecase) FetchByID(c context.Context, TaxSettingsID string) (domain.TaxSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.TaxSettingsRepository.FetchByID(ctx, TaxSettingsID)
}

func (lu *TaxSettingsUsecase) Fetch(c context.Context) ([]domain.TaxSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.TaxSettingsRepository.Fetch(ctx)
}
