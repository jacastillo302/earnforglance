package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/gdpr"
)

type gdprSettingsUsecase struct {
	GdprSettingsRepository domain.GdprSettingsRepository
	contextTimeout         time.Duration
}

func NewGdprSettingsUsecase(GdprSettingsRepository domain.GdprSettingsRepository, timeout time.Duration) domain.GdprSettingsRepository {
	return &gdprSettingsUsecase{
		GdprSettingsRepository: GdprSettingsRepository,
		contextTimeout:         timeout,
	}
}

func (tu *gdprSettingsUsecase) CreateMany(c context.Context, GdprSettingsList []domain.GdprSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.CreateMany(ctx, GdprSettingsList)
}

func (tu *gdprSettingsUsecase) Create(c context.Context, GdprSettings *domain.GdprSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Create(ctx, GdprSettings)
}

func (tu *gdprSettingsUsecase) Update(c context.Context, GdprSettings *domain.GdprSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Update(ctx, GdprSettings)
}

func (tu *gdprSettingsUsecase) Delete(c context.Context, GdprSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprSettingsRepository.Delete(ctx, GdprSettings)
}

func (lu *gdprSettingsUsecase) FetchByID(c context.Context, GdprSettingsID string) (domain.GdprSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprSettingsRepository.FetchByID(ctx, GdprSettingsID)
}

func (lu *gdprSettingsUsecase) Fetch(c context.Context) ([]domain.GdprSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprSettingsRepository.Fetch(ctx)
}
