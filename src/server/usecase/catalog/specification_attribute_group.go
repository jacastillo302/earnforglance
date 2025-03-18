package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type specificationattributegroupUsecase struct {
	specificationattributegroupRepository domain.SpecificationAttributeGroupRepository
	contextTimeout                        time.Duration
}

func NewSpecificationAttributeGroupUsecase(specificationattributegroupRepository domain.SpecificationAttributeGroupRepository, timeout time.Duration) domain.SpecificationAttributeGroupUsecase {
	return &specificationattributegroupUsecase{
		specificationattributegroupRepository: specificationattributegroupRepository,
		contextTimeout:                        timeout,
	}
}

func (tu *specificationattributegroupUsecase) Create(c context.Context, specificationattributegroup *domain.SpecificationAttributeGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributegroupRepository.Create(ctx, specificationattributegroup)
}

func (tu *specificationattributegroupUsecase) Update(c context.Context, specificationattributegroup *domain.SpecificationAttributeGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributegroupRepository.Update(ctx, specificationattributegroup)
}

func (tu *specificationattributegroupUsecase) Delete(c context.Context, specificationattributegroup *domain.SpecificationAttributeGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.specificationattributegroupRepository.Delete(ctx, specificationattributegroup)
}

func (lu *specificationattributegroupUsecase) FetchByID(c context.Context, specificationattributegroupID string) (domain.SpecificationAttributeGroup, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributegroupRepository.FetchByID(ctx, specificationattributegroupID)
}

func (lu *specificationattributegroupUsecase) Fetch(c context.Context) ([]domain.SpecificationAttributeGroup, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.specificationattributegroupRepository.Fetch(ctx)
}
