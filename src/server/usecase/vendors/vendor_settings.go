package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/vendors"
)

type vendorSettingsUsecase struct {
	vendorSettingsRepository domain.VendorSettingsRepository
	contextTimeout           time.Duration
}

func NewVendorSettingsUsecase(vendorSettingsRepository domain.VendorSettingsRepository, timeout time.Duration) domain.VendorSettingsUsecase {
	return &vendorSettingsUsecase{
		vendorSettingsRepository: vendorSettingsRepository,
		contextTimeout:           timeout,
	}
}

func (tu *vendorSettingsUsecase) Create(c context.Context, vendorSettings *domain.VendorSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorSettingsRepository.Create(ctx, vendorSettings)
}

func (tu *vendorSettingsUsecase) Update(c context.Context, vendorSettings *domain.VendorSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorSettingsRepository.Update(ctx, vendorSettings)
}

func (tu *vendorSettingsUsecase) Delete(c context.Context, vendorSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorSettingsRepository.Delete(ctx, vendorSettings)
}

func (lu *vendorSettingsUsecase) FetchByID(c context.Context, vendorSettingsID string) (domain.VendorSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorSettingsRepository.FetchByID(ctx, vendorSettingsID)
}

func (lu *vendorSettingsUsecase) Fetch(c context.Context) ([]domain.VendorSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorSettingsRepository.Fetch(ctx)
}
