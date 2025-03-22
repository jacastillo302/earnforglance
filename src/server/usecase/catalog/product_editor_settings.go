package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type producteditorsettingsUsecase struct {
	producteditorsettingsRepository domain.ProductEditorSettingsRepository
	contextTimeout                  time.Duration
}

func NewProductEditorSettingsUsecase(producteditorsettingsRepository domain.ProductEditorSettingsRepository, timeout time.Duration) domain.ProductEditorSettingsUsecase {
	return &producteditorsettingsUsecase{
		producteditorsettingsRepository: producteditorsettingsRepository,
		contextTimeout:                  timeout,
	}
}

func (tu *producteditorsettingsUsecase) Create(c context.Context, producteditorsettings *domain.ProductEditorSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producteditorsettingsRepository.Create(ctx, producteditorsettings)
}

func (tu *producteditorsettingsUsecase) Update(c context.Context, producteditorsettings *domain.ProductEditorSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producteditorsettingsRepository.Update(ctx, producteditorsettings)
}

func (tu *producteditorsettingsUsecase) Delete(c context.Context, producteditorsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.producteditorsettingsRepository.Delete(ctx, producteditorsettings)
}

func (lu *producteditorsettingsUsecase) FetchByID(c context.Context, producteditorsettingsID string) (domain.ProductEditorSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producteditorsettingsRepository.FetchByID(ctx, producteditorsettingsID)
}

func (lu *producteditorsettingsUsecase) Fetch(c context.Context) ([]domain.ProductEditorSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.producteditorsettingsRepository.Fetch(ctx)
}
