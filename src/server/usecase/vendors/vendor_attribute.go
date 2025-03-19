package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/vendors"
)

type vendorattributeUsecase struct {
	vendorattributeRepository domain.VendorAttributeRepository
	contextTimeout            time.Duration
}

func NewVendorAttributeUsecase(vendorattributeRepository domain.VendorAttributeRepository, timeout time.Duration) domain.VendorAttributeUsecase {
	return &vendorattributeUsecase{
		vendorattributeRepository: vendorattributeRepository,
		contextTimeout:            timeout,
	}
}

func (tu *vendorattributeUsecase) Create(c context.Context, vendorattribute *domain.VendorAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorattributeRepository.Create(ctx, vendorattribute)
}

func (tu *vendorattributeUsecase) Update(c context.Context, vendorattribute *domain.VendorAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorattributeRepository.Update(ctx, vendorattribute)
}

func (tu *vendorattributeUsecase) Delete(c context.Context, vendorattribute *domain.VendorAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorattributeRepository.Delete(ctx, vendorattribute)
}

func (lu *vendorattributeUsecase) FetchByID(c context.Context, vendorattributeID string) (domain.VendorAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorattributeRepository.FetchByID(ctx, vendorattributeID)
}

func (lu *vendorattributeUsecase) Fetch(c context.Context) ([]domain.VendorAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorattributeRepository.Fetch(ctx)
}
