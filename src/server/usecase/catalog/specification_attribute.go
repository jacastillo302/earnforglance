package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type specificationattributeUsecase struct {
	specificationattributeRepository domain.SpecificationAttributeRepository
	contextTimeout                   time.Duration
}

func NewSpecificationAttributeUsecase(specificationattributeRepository domain.SpecificationAttributeRepository, timeout time.Duration) domain.SpecificationAttributeUsecase {
	return &specificationattributeUsecase{
		specificationattributeRepository: specificationattributeRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *specificationattributeUsecase) Create(c context.Context, specificationattribute *domain.SpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeRepository.Create(ctx, specificationattribute)
}

func (tu *specificationattributeUsecase) Update(c context.Context, specificationattribute *domain.SpecificationAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeRepository.Update(ctx, specificationattribute)
}

func (tu *specificationattributeUsecase) Delete(c context.Context, specificationattribute string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributeRepository.Delete(ctx, specificationattribute)
}

func (lu *specificationattributeUsecase) FetchByID(c context.Context, specificationattributeID string) (domain.SpecificationAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributeRepository.FetchByID(ctx, specificationattributeID)
}

func (lu *specificationattributeUsecase) Fetch(c context.Context) ([]domain.SpecificationAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributeRepository.Fetch(ctx)
}
