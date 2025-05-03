package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/attributes"
)

type customerattributevalueUsecase struct {
	customerattributevalueRepository domain.PermisionRecordAttributeValueRepository
	contextTimeout                   time.Duration
}

func NewPermisionRecordAttributeValueUsecase(customerattributevalueRepository domain.PermisionRecordAttributeValueRepository, timeout time.Duration) domain.PermisionRecordAttributeValueUsecase {
	return &customerattributevalueUsecase{
		customerattributevalueRepository: customerattributevalueRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *customerattributevalueUsecase) CreateMany(c context.Context, items []domain.PermisionRecordAttributeValue) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributevalueRepository.CreateMany(ctx, items)
}

func (tu *customerattributevalueUsecase) Create(c context.Context, customerattributevalue *domain.PermisionRecordAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributevalueRepository.Create(ctx, customerattributevalue)
}

func (tu *customerattributevalueUsecase) Update(c context.Context, customerattributevalue *domain.PermisionRecordAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributevalueRepository.Update(ctx, customerattributevalue)
}

func (tu *customerattributevalueUsecase) Delete(c context.Context, customerattributevalue string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.customerattributevalueRepository.Delete(ctx, customerattributevalue)
}

func (lu *customerattributevalueUsecase) FetchByID(c context.Context, customerattributevalueID string) (domain.PermisionRecordAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerattributevalueRepository.FetchByID(ctx, customerattributevalueID)
}

func (lu *customerattributevalueUsecase) Fetch(c context.Context) ([]domain.PermisionRecordAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.customerattributevalueRepository.Fetch(ctx)
}
