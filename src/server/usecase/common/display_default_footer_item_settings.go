package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type displaydefaultfooteritemsettingsUsecase struct {
	displaydefaultfooteritemsettingsRepository domain.DisplayDefaultFooterItemSettingsRepository
	contextTimeout                             time.Duration
}

func NewDisplayDefaultFooterItemSettingsUsecase(displaydefaultfooteritemsettingsRepository domain.DisplayDefaultFooterItemSettingsRepository, timeout time.Duration) domain.DisplayDefaultFooterItemSettingsUsecase {
	return &displaydefaultfooteritemsettingsUsecase{
		displaydefaultfooteritemsettingsRepository: displaydefaultfooteritemsettingsRepository,
		contextTimeout: timeout,
	}
}

func (tu *displaydefaultfooteritemsettingsUsecase) CreateMany(c context.Context, items []domain.DisplayDefaultFooterItemSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultfooteritemsettingsRepository.CreateMany(ctx, items)
}

func (tu *displaydefaultfooteritemsettingsUsecase) Create(c context.Context, displaydefaultfooteritemsettings *domain.DisplayDefaultFooterItemSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultfooteritemsettingsRepository.Create(ctx, displaydefaultfooteritemsettings)
}

func (tu *displaydefaultfooteritemsettingsUsecase) Update(c context.Context, displaydefaultfooteritemsettings *domain.DisplayDefaultFooterItemSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultfooteritemsettingsRepository.Update(ctx, displaydefaultfooteritemsettings)
}

func (tu *displaydefaultfooteritemsettingsUsecase) Delete(c context.Context, displaydefaultfooteritemsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultfooteritemsettingsRepository.Delete(ctx, displaydefaultfooteritemsettings)
}

func (lu *displaydefaultfooteritemsettingsUsecase) FetchByID(c context.Context, displaydefaultfooteritemsettingsID string) (domain.DisplayDefaultFooterItemSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.displaydefaultfooteritemsettingsRepository.FetchByID(ctx, displaydefaultfooteritemsettingsID)
}

func (lu *displaydefaultfooteritemsettingsUsecase) Fetch(c context.Context) ([]domain.DisplayDefaultFooterItemSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.displaydefaultfooteritemsettingsRepository.Fetch(ctx)
}
