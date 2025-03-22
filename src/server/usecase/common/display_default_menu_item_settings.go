package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

// Compare this snippet from src/server/usecase/common/display_default_menu_item_settings.go:
type displaydefaultmenuitemsettingsUsecase struct {
	displaydefaultmenuitemsettingsRepository domain.DisplayDefaultMenuItemSettingsRepository
	contextTimeout                           time.Duration
}

func NewDisplayDefaultMenuItemSettingsUsecase(displaydefaultmenuitemsettingsRepository domain.DisplayDefaultMenuItemSettingsRepository, timeout time.Duration) domain.DisplayDefaultMenuItemSettingsUsecase {
	return &displaydefaultmenuitemsettingsUsecase{
		displaydefaultmenuitemsettingsRepository: displaydefaultmenuitemsettingsRepository,
		contextTimeout:                           timeout,
	}
}

func (tu *displaydefaultmenuitemsettingsUsecase) Create(c context.Context, displaydefaultmenuitemsettings *domain.DisplayDefaultMenuItemSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultmenuitemsettingsRepository.Create(ctx, displaydefaultmenuitemsettings)
}

func (tu *displaydefaultmenuitemsettingsUsecase) Update(c context.Context, displaydefaultmenuitemsettings *domain.DisplayDefaultMenuItemSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultmenuitemsettingsRepository.Update(ctx, displaydefaultmenuitemsettings)
}

func (tu *displaydefaultmenuitemsettingsUsecase) Delete(c context.Context, displaydefaultmenuitemsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.displaydefaultmenuitemsettingsRepository.Delete(ctx, displaydefaultmenuitemsettings)
}

func (lu *displaydefaultmenuitemsettingsUsecase) FetchByID(c context.Context, displaydefaultmenuitemsettingsID string) (domain.DisplayDefaultMenuItemSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.displaydefaultmenuitemsettingsRepository.FetchByID(ctx, displaydefaultmenuitemsettingsID)
}

func (lu *displaydefaultmenuitemsettingsUsecase) Fetch(c context.Context) ([]domain.DisplayDefaultMenuItemSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.displaydefaultmenuitemsettingsRepository.Fetch(ctx)
}
