package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/cms"
)

type widgetsettingsUsecase struct {
	widgetsettingsRepository domain.WidgetSettingsRepository
	contextTimeout           time.Duration
}

func NewWidgetSettingsUsecase(widgetsettingsRepository domain.WidgetSettingsRepository, timeout time.Duration) domain.WidgetSettingsUsecase {
	return &widgetsettingsUsecase{
		widgetsettingsRepository: widgetsettingsRepository,
		contextTimeout:           timeout,
	}
}

func (tu *widgetsettingsUsecase) CreateMany(c context.Context, items []domain.WidgetSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.widgetsettingsRepository.CreateMany(ctx, items)
}

func (tu *widgetsettingsUsecase) Create(c context.Context, widgetsettings *domain.WidgetSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.widgetsettingsRepository.Create(ctx, widgetsettings)
}

func (tu *widgetsettingsUsecase) Update(c context.Context, widgetsettings *domain.WidgetSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.widgetsettingsRepository.Update(ctx, widgetsettings)
}

func (tu *widgetsettingsUsecase) Delete(c context.Context, widgetsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.widgetsettingsRepository.Delete(ctx, widgetsettings)
}

func (lu *widgetsettingsUsecase) FetchByID(c context.Context, widgetsettingsID string) (domain.WidgetSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.widgetsettingsRepository.FetchByID(ctx, widgetsettingsID)
}

func (lu *widgetsettingsUsecase) Fetch(c context.Context) ([]domain.WidgetSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.widgetsettingsRepository.Fetch(ctx)
}
