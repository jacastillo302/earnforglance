package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type predefinedproductattributevalueUsecase struct {
	predefinedproductattributevalueRepository domain.PredefinedProductAttributeValueRepository
	contextTimeout                            time.Duration
}

func NewPredefinedProductAttributeValueUsecase(predefinedproductattributevalueRepository domain.PredefinedProductAttributeValueRepository, timeout time.Duration) domain.PredefinedProductAttributeValueUsecase {
	return &predefinedproductattributevalueUsecase{
		predefinedproductattributevalueRepository: predefinedproductattributevalueRepository,
		contextTimeout: timeout,
	}
}

func (tu *predefinedproductattributevalueUsecase) Create(c context.Context, predefinedproductattributevalue *domain.PredefinedProductAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.predefinedproductattributevalueRepository.Create(ctx, predefinedproductattributevalue)
}

func (tu *predefinedproductattributevalueUsecase) Update(c context.Context, predefinedproductattributevalue *domain.PredefinedProductAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.predefinedproductattributevalueRepository.Update(ctx, predefinedproductattributevalue)
}

func (tu *predefinedproductattributevalueUsecase) Delete(c context.Context, predefinedproductattributevalue string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.predefinedproductattributevalueRepository.Delete(ctx, predefinedproductattributevalue)
}

func (lu *predefinedproductattributevalueUsecase) FetchByID(c context.Context, predefinedproductattributevalueID string) (domain.PredefinedProductAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.predefinedproductattributevalueRepository.FetchByID(ctx, predefinedproductattributevalueID)
}

func (lu *predefinedproductattributevalueUsecase) Fetch(c context.Context) ([]domain.PredefinedProductAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.predefinedproductattributevalueRepository.Fetch(ctx)
}
