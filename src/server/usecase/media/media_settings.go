package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type mediaSettingsUsecase struct {
	mediaSettingsRepository domain.MediaSettingsRepository
	contextTimeout          time.Duration
}

func NewMediaSettingsUsecase(mediaSettingsRepository domain.MediaSettingsRepository, timeout time.Duration) domain.MediaSettingsUsecase {
	return &mediaSettingsUsecase{
		mediaSettingsRepository: mediaSettingsRepository,
		contextTimeout:          timeout,
	}
}

func (tu *mediaSettingsUsecase) CreateMany(c context.Context, items []domain.MediaSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.mediaSettingsRepository.CreateMany(ctx, items)
}

func (tu *mediaSettingsUsecase) Create(c context.Context, mediaSettings *domain.MediaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.mediaSettingsRepository.Create(ctx, mediaSettings)
}

func (tu *mediaSettingsUsecase) Update(c context.Context, mediaSettings *domain.MediaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.mediaSettingsRepository.Update(ctx, mediaSettings)
}

func (tu *mediaSettingsUsecase) Delete(c context.Context, mediaSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.mediaSettingsRepository.Delete(ctx, mediaSettings)
}

func (lu *mediaSettingsUsecase) FetchByID(c context.Context, mediaSettingsID string) (domain.MediaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.mediaSettingsRepository.FetchByID(ctx, mediaSettingsID)
}

func (lu *mediaSettingsUsecase) Fetch(c context.Context) ([]domain.MediaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.mediaSettingsRepository.Fetch(ctx)
}
