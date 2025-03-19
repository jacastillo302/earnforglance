package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/localization"
)

type localizationsettingsUsecase struct {
	localizationsettingsRepository domain.LocalizationSettingsRepository
	contextTimeout                 time.Duration
}

func NewLocalizationSettingsUsecase(localizationsettingsRepository domain.LocalizationSettingsRepository, timeout time.Duration) domain.LocalizationSettingsUsecase {
	return &localizationsettingsUsecase{
		localizationsettingsRepository: localizationsettingsRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *localizationsettingsUsecase) Create(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizationsettingsRepository.Create(ctx, localizationsettings)
}

func (tu *localizationsettingsUsecase) Update(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizationsettingsRepository.Update(ctx, localizationsettings)
}

func (tu *localizationsettingsUsecase) Delete(c context.Context, localizationsettings *domain.LocalizationSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizationsettingsRepository.Delete(ctx, localizationsettings)
}

func (lu *localizationsettingsUsecase) FetchByID(c context.Context, localizationsettingsID string) (domain.LocalizationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localizationsettingsRepository.FetchByID(ctx, localizationsettingsID)
}

func (lu *localizationsettingsUsecase) Fetch(c context.Context) ([]domain.LocalizationSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localizationsettingsRepository.Fetch(ctx)
}
