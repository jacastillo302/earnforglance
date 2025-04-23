package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type vendortUsecase struct {
	itemRepository domain.VendorRepository
	contextTimeout time.Duration
}

func NewVendortUsecase(itemRepository domain.VendorRepository, timeout time.Duration) domain.VendortUsecase {
	return &vendortUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *vendortUsecase) GetVendors(c context.Context, filter domain.VendorRequest) ([]domain.VendorsResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetVendors(ctx, filter)
}
