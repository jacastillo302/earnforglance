package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type currencysettingsUsecase struct {
	currencysettingsRepository domain.CurrencySettingsRepository
	contextTimeout             time.Duration
}

func NewCurrencySettingsUsecase(currencysettingsRepository domain.CurrencySettingsRepository, timeout time.Duration) domain.CurrencySettingsUsecase {
	return &currencysettingsUsecase{
		currencysettingsRepository: currencysettingsRepository,
		contextTimeout:             timeout,
	}
}

func (tu *currencysettingsUsecase) Create(c context.Context, currencysettings *domain.CurrencySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.currencysettingsRepository.Create(ctx, currencysettings)
}

func (tu *currencysettingsUsecase) Update(c context.Context, currencysettings *domain.CurrencySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.currencysettingsRepository.Update(ctx, currencysettings)
}

func (tu *currencysettingsUsecase) Delete(c context.Context, currencysettings *domain.CurrencySettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.currencysettingsRepository.Delete(ctx, currencysettings)
}

func (lu *currencysettingsUsecase) FetchByID(c context.Context, currencysettingsID string) (domain.CurrencySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.currencysettingsRepository.FetchByID(ctx, currencysettingsID)
}

func (lu *currencysettingsUsecase) Fetch(c context.Context) ([]domain.CurrencySettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.currencysettingsRepository.Fetch(ctx)
}
