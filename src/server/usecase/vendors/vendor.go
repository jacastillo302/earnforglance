package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/vendors"
)

type vendorUsecase struct {
	vendorRepository domain.VendorRepository
	contextTimeout   time.Duration
}

func NewVendorUsecase(vendorRepository domain.VendorRepository, timeout time.Duration) domain.VendorUsecase {
	return &vendorUsecase{
		vendorRepository: vendorRepository,
		contextTimeout:   timeout,
	}
}

func (tu *vendorUsecase) Create(c context.Context, vendor *domain.Vendor) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorRepository.Create(ctx, vendor)
}

func (tu *vendorUsecase) Update(c context.Context, vendor *domain.Vendor) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorRepository.Update(ctx, vendor)
}

func (tu *vendorUsecase) Delete(c context.Context, vendor string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorRepository.Delete(ctx, vendor)
}

func (lu *vendorUsecase) FetchByID(c context.Context, vendorID string) (domain.Vendor, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorRepository.FetchByID(ctx, vendorID)
}

func (lu *vendorUsecase) Fetch(c context.Context) ([]domain.Vendor, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorRepository.Fetch(ctx)
}
