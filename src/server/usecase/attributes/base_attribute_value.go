package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/attributes"
)

// baseattributevalueUsecase represents the usecase for the base attribute value
type baseattributevalueUsecase struct {
	baseattributevalueRepository domain.BaseAttributeValueRepository
	contextTimeout               time.Duration
}

func NewBaseAttributeValueUsecase(baseattributevalueRepository domain.BaseAttributeValueRepository, timeout time.Duration) domain.BaseAttributeValueUsecase {
	return &baseattributevalueUsecase{
		baseattributevalueRepository: baseattributevalueRepository,
		contextTimeout:               timeout,
	}
}

func (tu *baseattributevalueUsecase) Create(c context.Context, baseAttributeValue *domain.BaseAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributevalueRepository.Create(ctx, baseAttributeValue)
}

func (tu *baseattributevalueUsecase) Update(c context.Context, baseAttributeValue *domain.BaseAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributevalueRepository.Update(ctx, baseAttributeValue)
}

func (tu *baseattributevalueUsecase) Delete(c context.Context, baseAttributeValue *domain.BaseAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributevalueRepository.Delete(ctx, baseAttributeValue)
}

func (lu *baseattributevalueUsecase) FetchByID(c context.Context, baseAttributeValueID string) (domain.BaseAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.baseattributevalueRepository.FetchByID(ctx, baseAttributeValueID)
}

func (lu *baseattributevalueUsecase) Fetch(c context.Context) ([]domain.BaseAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.baseattributevalueRepository.Fetch(ctx)
}
