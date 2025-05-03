package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/attributes"
)

type customerattributeUsecase struct {
	customerattributeRepository domain.CustomerAttributeRepository
	contextTimeout              time.Duration
}

func NewCustomerAttributeUsecase(customerattributeRepository domain.CustomerAttributeRepository, timeout time.Duration) domain.CustomerAttributeUsecase {
	return &customerattributeUsecase{
		customerattributeRepository: customerattributeRepository,
		contextTimeout:              timeout,
	}
}

func (tu *customerattributeUsecase) CreateMany(c context.Context, items []domain.CustomerAttribute) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributeRepository.CreateMany(ctx, items)
}

func (tu *customerattributeUsecase) Create(c context.Context, customerattribute *domain.CustomerAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributeRepository.Create(ctx, customerattribute)
}

func (tu *customerattributeUsecase) Update(c context.Context, customerattribute *domain.CustomerAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributeRepository.Update(ctx, customerattribute)
}

func (tu *customerattributeUsecase) Delete(c context.Context, customerattribute string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributeRepository.Delete(ctx, customerattribute)
}

func (lu *customerattributeUsecase) FetchByID(c context.Context, customerattributeID string) (domain.CustomerAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerattributeRepository.FetchByID(ctx, customerattributeID)
}

func (lu *customerattributeUsecase) Fetch(c context.Context) ([]domain.CustomerAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerattributeRepository.Fetch(ctx)
}
