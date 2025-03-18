package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customerUsecase struct {
	customerRepository domain.CustomerRepository
	contextTimeout     time.Duration
}

func NewCustomerUsecase(customerRepository domain.CustomerRepository, timeout time.Duration) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
		contextTimeout:     timeout,
	}
}

func (tu *customerUsecase) Create(c context.Context, customer *domain.Customer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRepository.Create(ctx, customer)
}

func (tu *customerUsecase) Update(c context.Context, customer *domain.Customer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRepository.Update(ctx, customer)
}

func (tu *customerUsecase) Delete(c context.Context, customer *domain.Customer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerRepository.Delete(ctx, customer)
}

func (lu *customerUsecase) FetchByID(c context.Context, customerID string) (domain.Customer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerRepository.FetchByID(ctx, customerID)
}

func (lu *customerUsecase) Fetch(c context.Context) ([]domain.Customer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerRepository.Fetch(ctx)
}
