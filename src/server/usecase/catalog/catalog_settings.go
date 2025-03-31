package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type catalogsettingsUsecase struct {
	catalogsettingsRepository domain.CatalogSettingsRepository
	contextTimeout            time.Duration
}

func NewCatalogSettingsUsecase(catalogsettingsRepository domain.CatalogSettingsRepository, timeout time.Duration) domain.CatalogSettingsUsecase {
	return &catalogsettingsUsecase{
		catalogsettingsRepository: catalogsettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *catalogsettingsUsecase) CreateMany(c context.Context, items []domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.CreateMany(ctx, items)
}

func (tu *catalogsettingsUsecase) Create(c context.Context, item *domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Create(ctx, item)
}

func (tu *catalogsettingsUsecase) Update(c context.Context, item *domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Update(ctx, item)
}

func (tu *catalogsettingsUsecase) Delete(c context.Context, item string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Delete(ctx, item)
}

func (lu *catalogsettingsUsecase) FetchByID(c context.Context, itemID string) (domain.CatalogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.catalogsettingsRepository.FetchByID(ctx, itemID)
}

func (lu *catalogsettingsUsecase) Fetch(c context.Context) ([]domain.CatalogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.catalogsettingsRepository.Fetch(ctx)
}
