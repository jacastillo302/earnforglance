package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customerCustomerRoleMappingUsecase struct {
	customerCustomerRoleMappingRepository domain.CustomerCustomerRoleMappingRepository
	contextTimeout                        time.Duration
}

func NewCustomerCustomerRoleMappingUsecase(customerCustomerRoleMappingRepository domain.CustomerCustomerRoleMappingRepository, timeout time.Duration) domain.CustomerCustomerRoleMappingUsecase {
	return &customerCustomerRoleMappingUsecase{
		customerCustomerRoleMappingRepository: customerCustomerRoleMappingRepository,
		contextTimeout:                        timeout,
	}
}

func (tu *customerCustomerRoleMappingUsecase) Create(c context.Context, customerCustomerRoleMapping *domain.CustomerCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerCustomerRoleMappingRepository.Create(ctx, customerCustomerRoleMapping)
}

func (tu *customerCustomerRoleMappingUsecase) Update(c context.Context, customerCustomerRoleMapping *domain.CustomerCustomerRoleMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerCustomerRoleMappingRepository.Update(ctx, customerCustomerRoleMapping)
}

func (tu *customerCustomerRoleMappingUsecase) Delete(c context.Context, customerCustomerRoleMapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerCustomerRoleMappingRepository.Delete(ctx, customerCustomerRoleMapping)
}

func (lu *customerCustomerRoleMappingUsecase) FetchByID(c context.Context, customerCustomerRoleMappingID string) (domain.CustomerCustomerRoleMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerCustomerRoleMappingRepository.FetchByID(ctx, customerCustomerRoleMappingID)
}

func (lu *customerCustomerRoleMappingUsecase) Fetch(c context.Context) ([]domain.CustomerCustomerRoleMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerCustomerRoleMappingRepository.Fetch(ctx)
}
