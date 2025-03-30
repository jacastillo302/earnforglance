package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type customerpasswordUsecase struct {
	customerpasswordRepository domain.CustomerPasswordRepository
	contextTimeout             time.Duration
}

func NewCustomerPasswordUsecase(customerpasswordRepository domain.CustomerPasswordRepository, timeout time.Duration) domain.CustomerPasswordUsecase {
	return &customerpasswordUsecase{
		customerpasswordRepository: customerpasswordRepository,
		contextTimeout:             timeout,
	}
}

func (tu *customerpasswordUsecase) CreateMany(c context.Context, items []domain.CustomerPassword) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerpasswordRepository.CreateMany(ctx, items)
}

func (tu *customerpasswordUsecase) Create(c context.Context, customerpassword *domain.CustomerPassword) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerpasswordRepository.Create(ctx, customerpassword)
}

func (tu *customerpasswordUsecase) Update(c context.Context, customerpassword *domain.CustomerPassword) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerpasswordRepository.Update(ctx, customerpassword)
}

func (tu *customerpasswordUsecase) Delete(c context.Context, customerpassword string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerpasswordRepository.Delete(ctx, customerpassword)
}

func (lu *customerpasswordUsecase) FetchByID(c context.Context, customerpasswordID string) (domain.CustomerPassword, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerpasswordRepository.FetchByID(ctx, customerpasswordID)
}

func (lu *customerpasswordUsecase) Fetch(c context.Context) ([]domain.CustomerPassword, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerpasswordRepository.Fetch(ctx)
}
