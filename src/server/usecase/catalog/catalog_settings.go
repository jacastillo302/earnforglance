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

func (tu *catalogsettingsUsecase) Create(c context.Context, affiliate *domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Create(ctx, affiliate)
}

func (tu *catalogsettingsUsecase) Update(c context.Context, affiliate *domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Update(ctx, affiliate)
}

func (tu *catalogsettingsUsecase) Delete(c context.Context, affiliate *domain.CatalogSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.catalogsettingsRepository.Delete(ctx, affiliate)
}

func (lu *catalogsettingsUsecase) FetchByID(c context.Context, affiliateID string) (domain.CatalogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.catalogsettingsRepository.FetchByID(ctx, affiliateID)
}

func (lu *catalogsettingsUsecase) Fetch(c context.Context) ([]domain.CatalogSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.catalogsettingsRepository.Fetch(ctx)
}
