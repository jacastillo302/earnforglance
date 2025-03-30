package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customerRoleUsecase struct {
	customerRoleRepository domain.CustomerRoleRepository
	contextTimeout         time.Duration
}

func NewCustomerRoleUsecase(customerRoleRepository domain.CustomerRoleRepository, timeout time.Duration) domain.CustomerRoleUsecase {
	return &customerRoleUsecase{
		customerRoleRepository: customerRoleRepository,
		contextTimeout:         timeout,
	}
}

func (tu *customerRoleUsecase) CreateMany(c context.Context, items []domain.CustomerRole) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRoleRepository.CreateMany(ctx, items)
}

func (tu *customerRoleUsecase) Create(c context.Context, customerRole *domain.CustomerRole) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRoleRepository.Create(ctx, customerRole)
}

func (tu *customerRoleUsecase) Update(c context.Context, customerRole *domain.CustomerRole) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRoleRepository.Update(ctx, customerRole)
}

func (tu *customerRoleUsecase) Delete(c context.Context, customerRole string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRoleRepository.Delete(ctx, customerRole)
}

func (lu *customerRoleUsecase) FetchByID(c context.Context, customerRoleID string) (domain.CustomerRole, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerRoleRepository.FetchByID(ctx, customerRoleID)
}

func (lu *customerRoleUsecase) Fetch(c context.Context) ([]domain.CustomerRole, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerRoleRepository.Fetch(ctx)
}
