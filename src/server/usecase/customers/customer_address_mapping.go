package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customeraddressmappingUsecase struct {
	customeraddressmappingRepository domain.CustomerAddressMappingRepository
	contextTimeout                   time.Duration
}

func NewCustomerAddressMappingUsecase(customeraddressmappingRepository domain.CustomerAddressMappingRepository, timeout time.Duration) domain.CustomerAddressMappingUsecase {
	return &customeraddressmappingUsecase{
		customeraddressmappingRepository: customeraddressmappingRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *customeraddressmappingUsecase) Create(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customeraddressmappingRepository.Create(ctx, customeraddressmapping)
}

func (tu *customeraddressmappingUsecase) Update(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customeraddressmappingRepository.Update(ctx, customeraddressmapping)
}

func (tu *customeraddressmappingUsecase) Delete(c context.Context, customeraddressmapping *domain.CustomerAddressMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customeraddressmappingRepository.Delete(ctx, customeraddressmapping)
}

func (lu *customeraddressmappingUsecase) FetchByID(c context.Context, customeraddressmappingID string) (domain.CustomerAddressMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customeraddressmappingRepository.FetchByID(ctx, customeraddressmappingID)
}

func (lu *customeraddressmappingUsecase) Fetch(c context.Context) ([]domain.CustomerAddressMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customeraddressmappingRepository.Fetch(ctx)
}
