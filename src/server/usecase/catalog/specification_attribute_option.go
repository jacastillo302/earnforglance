package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type specificationattributeoptionUsecase struct {
	specificationattributeoptionRepository domain.SpecificationAttributeOptionRepository
	contextTimeout                         time.Duration
}

func NewSpecificationAttributeOptionUsecase(specificationattributeoptionRepository domain.SpecificationAttributeOptionRepository, timeout time.Duration) domain.SpecificationAttributeOptionUsecase {
	return &specificationattributeoptionUsecase{
		specificationattributeoptionRepository: specificationattributeoptionRepository,
		contextTimeout:                         timeout,
	}
}

func (tu *specificationattributeoptionUsecase) Create(c context.Context, specificationattributeoption *domain.SpecificationAttributeOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeoptionRepository.Create(ctx, specificationattributeoption)
}

func (tu *specificationattributeoptionUsecase) Update(c context.Context, specificationattributeoption *domain.SpecificationAttributeOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeoptionRepository.Update(ctx, specificationattributeoption)
}

func (tu *specificationattributeoptionUsecase) Delete(c context.Context, specificationattributeoption *domain.SpecificationAttributeOption) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeoptionRepository.Delete(ctx, specificationattributeoption)
}

func (lu *specificationattributeoptionUsecase) FetchByID(c context.Context, specificationattributeoptionID string) (domain.SpecificationAttributeOption, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributeoptionRepository.FetchByID(ctx, specificationattributeoptionID)
}

func (lu *specificationattributeoptionUsecase) Fetch(c context.Context) ([]domain.SpecificationAttributeOption, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributeoptionRepository.Fetch(ctx)
}
