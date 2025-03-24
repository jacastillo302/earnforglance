package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type measuresettingsUsecase struct {
	measuresettingsRepository domain.MeasureSettingsRepository
	contextTimeout            time.Duration
}

func NewMeasureSettingsUsecase(measuresettingsRepository domain.MeasureSettingsRepository, timeout time.Duration) domain.MeasureSettingsUsecase {
	return &measuresettingsUsecase{
		measuresettingsRepository: measuresettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *measuresettingsUsecase) CreateMany(c context.Context, items []domain.MeasureSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measuresettingsRepository.CreateMany(ctx, items)
}

func (tu *measuresettingsUsecase) Create(c context.Context, measuresettings *domain.MeasureSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measuresettingsRepository.Create(ctx, measuresettings)
}

func (tu *measuresettingsUsecase) Update(c context.Context, measuresettings *domain.MeasureSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measuresettingsRepository.Update(ctx, measuresettings)
}

func (tu *measuresettingsUsecase) Delete(c context.Context, measuresettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measuresettingsRepository.Delete(ctx, measuresettings)
}

func (lu *measuresettingsUsecase) FetchByID(c context.Context, measuresettingsID string) (domain.MeasureSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.measuresettingsRepository.FetchByID(ctx, measuresettingsID)
}

func (lu *measuresettingsUsecase) Fetch(c context.Context) ([]domain.MeasureSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.measuresettingsRepository.Fetch(ctx)
}
