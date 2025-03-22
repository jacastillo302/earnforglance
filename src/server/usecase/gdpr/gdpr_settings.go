package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/gdpr"
)

type GdprSettingsUsecase struct {
	GdprSettingsRepository domain.GdprSettingsRepository
	contextTimeout         time.Duration
}

func NewGdprSettingsUsecase(GdprSettingsRepository domain.GdprSettingsRepository, timeout time.Duration) domain.GdprSettingsRepository {
	return &GdprSettingsUsecase{
		GdprSettingsRepository: GdprSettingsRepository,
		contextTimeout:         timeout,
	}
}

func (tu *GdprSettingsUsecase) Create(c context.Context, GdprSettings *domain.GdprSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Create(ctx, GdprSettings)
}

func (tu *GdprSettingsUsecase) Update(c context.Context, GdprSettings *domain.GdprSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Update(ctx, GdprSettings)
}

func (tu *GdprSettingsUsecase) Delete(c context.Context, GdprSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Delete(ctx, GdprSettings)
}

func (lu *GdprSettingsUsecase) FetchByID(c context.Context, GdprSettingsID string) (domain.GdprSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprSettingsRepository.FetchByID(ctx, GdprSettingsID)
}

func (lu *GdprSettingsUsecase) Fetch(c context.Context) ([]domain.GdprSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprSettingsRepository.Fetch(ctx)
}
