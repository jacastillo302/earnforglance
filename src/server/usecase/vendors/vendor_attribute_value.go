package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/vendors"
)

type vendorAttributeValueUsecase struct {
	vendorAttributeValueRepository domain.VendorAttributeValueRepository
	contextTimeout                 time.Duration
}

func NewVendorAttributeValueUsecase(vendorAttributeValueRepository domain.VendorAttributeValueRepository, timeout time.Duration) domain.VendorAttributeValueUsecase {
	return &vendorAttributeValueUsecase{
		vendorAttributeValueRepository: vendorAttributeValueRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *vendorAttributeValueUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.VendorAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorAttributeValueRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *vendorAttributeValueUsecase) Create(c context.Context, vendorAttributeValue *domain.VendorAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorAttributeValueRepository.Create(ctx, vendorAttributeValue)
}

func (tu *vendorAttributeValueUsecase) Update(c context.Context, vendorAttributeValue *domain.VendorAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorAttributeValueRepository.Update(ctx, vendorAttributeValue)
}

func (tu *vendorAttributeValueUsecase) Delete(c context.Context, vendorAttributeValue string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.vendorAttributeValueRepository.Delete(ctx, vendorAttributeValue)
}

func (lu *vendorAttributeValueUsecase) FetchByID(c context.Context, vendorAttributeValueID string) (domain.VendorAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorAttributeValueRepository.FetchByID(ctx, vendorAttributeValueID)
}

func (lu *vendorAttributeValueUsecase) Fetch(c context.Context) ([]domain.VendorAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.vendorAttributeValueRepository.Fetch(ctx)
}
